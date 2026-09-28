package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("no website provided\n")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Printf("too many arguments provided\n")
		os.Exit(1)
	}

	fmt.Printf("starting crawler of: %v\n", os.Args[1])

	rawUrl := os.Args[1]

	const maxConcurrency = 10
	cfg, err := configure(rawUrl, maxConcurrency)
	if err != nil {
		fmt.Printf("error making new config struct: %v", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawUrl)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
