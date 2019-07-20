package sitemap

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

const (
	header      = `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
	footer      = `</urlset>`
	indexHeader = `<?xml version="1.0" encoding="UTF-8"?><sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
	indexFooter = `</sitemapindex>`
	tmpExt      = `.xml.tmp`
	ext         = `.xml`
)

// Options is a struct for specifying configuration options for the sitemap.Generator object.
type Options struct {
	// Filename is base file name for sitemap w/o extension
	Filename string
	// Max file size (default 10485760)
	MaxFileSize int
	// Max links in one file (default 50000)
	MaxURLs int
	// Dir keeps directory name for sitemap files
	Dir string
	// BaseURL used for generate sitemap index file
	BaseURL string
}

// New constructs a new Generator instance with the supplied options.
func New(options ...Options) *Generator {
	var o Options
	if len(options) == 0 {
		o = Options{}
	} else {
		o = options[0]
	}
	g := Generator{
		opt: o,
	}
	return &g
}

// Generator is a service that provides functions for creating sitemap files
type Generator struct {
	// opt keeps generator options
	opt Options
	// fileCount keeps number of completed files
	fileCount int
	// body keeps current file body length
	bodyLen int
	// bodyURLCount
	bodyURLCount int
	// maxBodyLen keeps max allowed body len
	maxBodyLen int
	// temp file
	tmpFile *os.File
}

// Open drops internal counters and checks you generator options
func (g *Generator) Open() (err error) {
	if g.opt.MaxFileSize == 0 {
		g.opt.MaxFileSize = 10485760
	}
	g.maxBodyLen = g.opt.MaxFileSize - len(header) - len(footer)
	if g.maxBodyLen <= 0 {
		return errors.New("invalid MaxFileSize option value")
	}
	if g.opt.MaxURLs == 0 {
		g.opt.MaxURLs = 50000
	}
	if g.opt.MaxURLs <= 0 {
		return errors.New("invalid MaxURLs option value")
	}
	if g.opt.Filename == "" {
		g.opt.Filename = "sitemap"
	}
	if g.opt.Dir, err = filepath.Abs(g.opt.Dir); err != nil {
		return err
	}
	if err = os.MkdirAll(g.opt.Dir, os.ModePerm); err != nil {
		return err
	}
	if g.opt.BaseURL == "" {
		return errors.New("empty BaseURL option value")
	}
	g.fileCount = 0
	return g.createTmp()
}

// Close finishes generation process, all temp files will removed
func (g *Generator) Close() error {
	if g.fileCount > 0 {
		g.fileCount++ // add sitemaindex
	}
	if err := g.closeAndRenameTmp(); err != nil {
		return err
	}
	if g.fileCount > 0 {
		g.createIndex()
	}
	for i := 0; i <= g.fileCount; i++ {
		if err := os.Rename(g.filename(i, tmpExt), g.filename(i, ext)); err != nil {
			return err
		}
	}
	return nil
}

// Abort stops generator and clears all temporary files
func (g *Generator) Abort() error {
	for i := 0; i <= g.fileCount; i++ {
		if err := os.Remove(g.filename(i, tmpExt)); err != nil {
			return err
		}
	}
	return nil
}

// Add new url to sitemap
func (g *Generator) Add(url SitemapURL) error {
	node := g.formatURLNode(url)
	if !g.canFit(len(node)) {
		if err := g.nextTmp(); err != nil {
			return err
		}
	}
	l, err := g.tmpFile.Write([]byte(node))
	g.bodyLen += l
	g.bodyURLCount++
	return err
}

// canFit returns true if one more node with specified length can be added to file
func (g *Generator) canFit(nodeLen int) bool {
	return g.bodyURLCount < g.opt.MaxURLs && g.bodyLen+nodeLen <= g.maxBodyLen
}

// formatURLNode creates xml for url node
func (g *Generator) formatURLNode(u SitemapURL) string {
	r := `<url><loc>` + u.SitemapLoc() + `</loc>`
	if t := u.SitemapLastMod(); t != "" {
		r += `<lastmod>` + t + `</lastmod>`
	}
	if t := u.SitemapChangeFreq(); t != "" {
		r += `<changefreq>` + t + `</changefreq>`
	}
	if t := u.SitemapPriority(); t != "" {
		r += `<priority>` + t + `</priority>`
	}
	return r + `</url>`
}

// newTmpFiles creates new tmp file of files (if gzip required)
// and returns slice of file descriptors
func (g *Generator) createTmp() (err error) {
	g.tmpFile, err = os.Create(filepath.Join(g.opt.Dir, g.opt.Filename+tmpExt))
	g.bodyLen = 0
	g.bodyURLCount = 0
	if err != nil {
		return err
	}
	_, err = g.tmpFile.Write([]byte(header))
	return err
}

// newTmpFiles creates new tmp file of files (if gzip required)
// and returns slice of file descriptors
func (g *Generator) createIndex() error {
	f, err := os.Create(g.filename(0, tmpExt))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(indexHeader))
	if err != nil {
		return err
	}
	for i := 0; i < g.fileCount; i++ {
		_, err = f.Write([]byte(`<sitemap><loc>` + g.opt.BaseURL + g.opt.Filename + `-` +
			strconv.Itoa(i+1) +
			ext + `</loc></sitemap>`))
	}
	_, err = f.Write([]byte(indexFooter))
	if err != nil {
		return err
	}
	return err
}

// creates new temp files
func (g *Generator) nextTmp() error {
	g.fileCount++
	if err := g.closeAndRenameTmp(); err != nil {
		return err
	}
	return g.createTmp()
}

// closeAndRenameTmp closes and renames tmp file
// sitemap.xml.tmp => sitemap-1.xml.tmp
func (g *Generator) closeAndRenameTmp() error {
	if g.tmpFile == nil {
		return errors.New("trying to close empty tmp file")
	}
	if _, err := g.tmpFile.Write([]byte(footer)); err != nil {
		return err
	}
	tmpPath := g.tmpFile.Name()
	err := g.tmpFile.Close()
	if err != nil {
		return err
	}
	if g.fileCount > 0 {
		return os.Rename(tmpPath, g.filename(g.fileCount, tmpExt))
	}
	return nil
}

// filename returns filename based on config index and specified extension
func (g *Generator) filename(i int, ext string) string {
	s := ""
	if i > 0 {
		s = "-" + strconv.Itoa(i)
	}
	return filepath.Join(g.opt.Dir, g.opt.Filename+s+ext)
}
