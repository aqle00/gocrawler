package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 4 {
		fmt.Printf("too many arguments provided\n")
		os.Exit(1)
	}

	if len(os.Args) != 4 {
		fmt.Printf("not enough arguments\n")
		fmt.Printf("command usage: crawler website maxConcurrency maxPage\n")
		os.Exit(1)
	}

	fmt.Printf("starting crawler of: %v\n", os.Args[1])
	rawUrl := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("invalid command: maxConcurrency requires a number\n")
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("invalid command: maxPage requires a number\n")
	}

	cfg, err := configure(rawUrl, maxConcurrency, maxPages)
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
