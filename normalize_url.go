package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", err)
	}

	normalURL := parsedURL.Host + parsedURL.Path
	normalURL = strings.ToLower(strings.TrimSuffix(normalURL, "/"))

	return normalURL, nil
}
