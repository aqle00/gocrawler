package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	//send empty struct to concurrency control to block if buffer size reaches max
	// this is to prevent the crawler from going crazy spawning too many goroutines
	// if too many goroutines -> too many http requests-> server will ban IP cuz we spamming
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.overCrawlLimit() {
		return
	}

	// check if rawCurrentURL belongs to baseURL
	// if not just return early
	//example: currentURL: fb.com/something, baseURL: fb.com ->crawl
	//						google.com/somthgn			fb.com -> return
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldnt parse base url, error: %v\n", err)
		return
	}
	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		fmt.Printf("url '%v' outside of '%v'\n", rawCurrentURL, cfg.baseURL)
		return
	}

	// get normalizedUrl to check if page has been visited or not
	normalizedUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("invalid url, error: %v", err)
		return
	}

	//check if page has already been visited
	// !isFirst means page already visited, return early if so
	isFirst := cfg.addPageVisit(normalizedUrl)
	if !isFirst {
		return
	}

	// reachest here, page not visited, start crawling logic
	fmt.Printf("crawling url '%v'\n", rawCurrentURL)
	//extract html for crawl
	res, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldnt get HTML, error: %v\n", err)
		return
	}

	//crawl and extract stuff from the html
	pageData := extractPageData(res, rawCurrentURL)
	cfg.setPageData(normalizedUrl, pageData)

	// recursively crawl found URLs
	// fmt.Printf("found urls: '%v'\n", foundURLs)
	for _, url := range pageData.OutgoingLinks {
		fmt.Printf("recursively crawling url '%v'", url)
		cfg.wg.Add(1)
		cfg.wg.Go(func() {
			cfg.crawlPage(url)
		})
	}
}
