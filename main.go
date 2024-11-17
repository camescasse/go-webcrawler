package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	pages := map[string]int{}
	crawlPage(rawBaseURL, rawBaseURL, pages)

	fmt.Println()
	fmt.Println("results:")
	for page, count := range pages {
		fmt.Printf("%s: %d\n", page, count)
	}

	os.Exit(0)
}
