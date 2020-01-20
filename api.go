package et

import (
	"fmt"
	"strings"

	"crawler.club/dl"
)

func (p *Parser) ParseURL(url string) ([]*UrlTask, []map[string]interface{}, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return nil, nil, fmt.Errorf("empty url")
	}
	req := &dl.HttpRequest{Url: url}
	if p.UA == "google" || p.UA == "pc" || p.UA == "mobile" {
		req.Platform = p.UA
	}
	resp := dl.Download(req)
	if resp.Error != nil {
		return nil, nil, resp.Error
	}
	return p.Parse(resp.Text, url)
}

func (p *Parser) Do() ([]*UrlTask, []map[string]interface{}, error) {
	return p.ParseURL(p.ExampleUrl)
}
