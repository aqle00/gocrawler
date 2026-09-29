package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)

	baseUrl, err := url.Parse(pageURL)
	if err != nil {
		fmt.Printf("couldn't parse page url: %v, error: %v\n", pageURL, err)
		return PageData{
			URL:            pageURL,
			Heading:        heading,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	urls, err := getURLsFromHTML(html, baseUrl)
	if err != nil {
		fmt.Printf("couldn't get urls, error: %v\n", err)
		urls = nil
	}
	images, err := getImagesFromHTML(html, baseUrl)
	if err != nil {
		fmt.Printf("couldn't get images, error: %v\n", err)
		images = nil
	}

	return PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  urls,
		ImageURLs:      images,
	}
}
