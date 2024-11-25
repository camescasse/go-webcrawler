package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]
	maxConcurrency := 10

	config, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Println("error starting config: %w", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	config.wg.Add(1)
	go config.crawlPage(rawBaseURL)
	config.wg.Wait()

	fmt.Println()
	fmt.Println("results:")
	for page, count := range config.pages {
		fmt.Printf("%s: %d\n", page, count)
	}
}
