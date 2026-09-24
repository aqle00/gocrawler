package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("could not parse url: %v", err)
	}

	var foundURLs []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok || strings.TrimSpace(url) == "" {
			return
		}

		parsedURL, err := baseURL.Parse(url)
		if err != nil {
			fmt.Printf("could not parse url: %v, err: %v", url, err)
			return
		}
		foundURLs = append(foundURLs, parsedURL.String())
	})

	return foundURLs, nil
}
