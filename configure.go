package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	maxPages           int
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, exist := cfg.pages[normalizedURL]; !exist {
		page := PageData{
			URL: normalizedURL,
		}
		cfg.pages[normalizedURL] = page
		return true
	}
	return false
}

func (cfg *config) setPageData(normalizedURL string, pageData PageData) {
	// make new pageData and return, locking beefore and unlocking after
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = pageData
}

func configure(rawBaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	// make a new config struct with these data
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldnt parse base URL: %v", err)
	}

	return &config{
		maxPages:           maxPages,
		pages:              make(map[string]PageData),
		baseURL:            baseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}

func (cfg *config) overCrawlLimit() (overcap bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if len(cfg.pages) > cfg.maxPages {
		return true
	}
	return false
}

// func (cfg *config) pagesLen() int {
// 	cfg.mu.Lock()
// 	defer cfg.mu.Unlock()
// 	return len(cfg.pages)
// }
