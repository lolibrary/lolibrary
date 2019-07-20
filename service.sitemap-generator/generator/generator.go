package generator

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.sitemap-generator/dao"
	"github.com/lolibrary/lolibrary/service.sitemap-generator/domain"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/pahanini/go-sitemap-generator"
)

var SitemapDir = path.Join(os.TempDir(), "sitemap")

func init() {
	override := os.Getenv("SITEMAP_DIR")
	if override != "" {
		SitemapDir = override
	}
}

func Generate(ctx context.Context) error {
	sm := sitemap.New(sitemap.Options{
		Dir: SitemapDir,
		BaseURL: "https://lolibrary.org",
		MaxURLs: 15000,
		MaxFileSize: 10*1024*1024,
	})

	if err := sm.Open(); err != nil {
		slog.Error(ctx, "Error opening sitemap filestream: %v", err)
		return terrors.Wrap(err, nil)
	}

	if err := addStaticURLs(sm); err != nil {
		slog.Error(ctx, "Error adding static URLs: %v", err)
		_ = sm.Abort()
		return terrors.Wrap(err, nil)
	}

	if err := dao.SitemapItemCursor(ctx, func(item *domain.SitemapItem) error {
		u := sitemap.URL{
			Loc:        fmt.Sprintf("/items/%s", item.Slug),
			LastMod:    util.TimeToRFC3339(item.UpdatedAt),
		}
		if err := sm.Add(u); err != nil {
			slog.Error(ctx, "Error adding sitemap item to index: %v", err, map[string]string{
				"slug": item.Slug,
			})
			_ = sm.Abort()
			return err
		}

		return nil
	}); err != nil {
		_ = sm.Abort()
		return terrors.Wrap(err, nil)
	}

	if err := sm.Close(); err != nil {
		slog.Error(ctx, "Error closing sitemap (rename + replace files): %v", err)
		return terrors.Wrap(err, nil)
	}

	return nil
}

func addStaticURLs(gen *sitemap.Generator) error {
	urls := [...]string{"/", "/donate", "/search", "/login", "/register"}

	for _, u := range urls {
		if err := gen.Add(sitemap.URL{Loc: u}); err != nil {
			return err
		}
	}

	return nil
}