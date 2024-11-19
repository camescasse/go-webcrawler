package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("error parsing argument:", err)
		os.Exit(1)
	}

	config := config{
		pages:   make(map[string]int),
		baseURL: rawBaseURL,
	}

	fmt.Printf("starting crawl of: %s...\n", config.baseURL.String())
	config.crawlPage(config.baseURL)

	fmt.Println()
	fmt.Println("results:")
	for page, count := range config.pages {
		fmt.Printf("%s: %d\n", page, count)
	}

	os.Exit(0)
}
