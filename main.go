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
	pages := map[string]int{}
	fmt.Println("starting crawl...")
	result, err := crawlPage(rawBaseURL, rawBaseURL, pages)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println("results:")
	for k, v := range result {
		fmt.Printf("%s: %d\n", k, v)
	}

	os.Exit(0)
}
