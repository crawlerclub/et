package et

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type parsers struct {
	sync.Mutex
	items map[string]*Parser
}

func (p *parsers) get(fname string, refresh bool) (*Parser, error) {
	p.Lock()
	defer p.Unlock()
	if !refresh && p.items[fname] != nil {
		return p.items[fname], nil
	}
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	parser := new(Parser)
	if err := json.Unmarshal(content, parser); err != nil {
		return nil, err
	}
	p.items[fname] = parser
	return parser, nil
}

var pool = &parsers{items: make(map[string]*Parser)}

func Parse(fname, url, page string) ([]map[string]interface{}, error) {
	p, err := pool.get(fname, false)
	if err != nil {
		return nil, err
	}
	_, ret, err := p.Parse(page, url)
	return ret, err
}

func ParseExt(fname, url, page string) (string, error) {
	ret, err := Parse(fname, url, page)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(ret)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
