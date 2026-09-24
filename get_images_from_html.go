package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("could not parse image: %v", err)
	}

	var foundImages []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		image, ok := s.Attr("src")
		if !ok || strings.TrimSpace(image) == "" {
			return
		}

		parsedURL, err := baseURL.Parse(image)
		if err != nil {
			fmt.Printf("could not parse image src: %v, err: %v", image, err)
			return
		}
		foundImages = append(foundImages, parsedURL.String())
	})

	return foundImages, nil
}
