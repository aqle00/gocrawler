package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("couldnt create new request, error: %v\n", err)
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http request failed, error: %v\n", err)
	}
	defer res.Body.Close()

	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("response format is not text/html\n")
	}

	if res.StatusCode > 399 {
		return "", fmt.Errorf("http request failed with code: %v\n", res.StatusCode)
	}

	htmlResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body, error: %v\n", err)
	}

	return string(htmlResponse), nil
}
