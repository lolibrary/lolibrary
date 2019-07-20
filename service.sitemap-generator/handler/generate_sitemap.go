package handler

import (
	"github.com/lolibrary/lolibrary/service.sitemap-generator/generator"
	sitemapgeneratorproto "github.com/lolibrary/lolibrary/service.sitemap-generator/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleGenerateSitemap(req typhon.Request) typhon.Response {
	body := &sitemapgeneratorproto.POSTGenerateSitemapRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}
	if err := generator.Generate(req); err != nil {
		slog.Error(req, "Error generating sitemap: %v", err)
		return typhon.Response{Error: err}
	}


	return req.Response(&sitemapgeneratorproto.POSTGenerateSitemapResponse{})
}
