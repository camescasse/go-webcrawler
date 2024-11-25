package main

import (
	"fmt"
	"net/url"
)

func (config *config) crawlPage(rawCurrentURL string) {
	config.concurrencyControl <- struct{}{}
	defer func() {
		<-config.concurrencyControl
		config.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing raw url '%s': %v\n", rawCurrentURL, err)
		return
	}

	if config.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(currentURL.String())
	if err != nil {
		fmt.Println("error normalizing current url:", err)
		return
	}

	isFirst := config.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Println("crawling:", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting page html for '%s': %v\n", rawCurrentURL, err)
		return
	}

	foundURLs, err := getURLsFromHTML(html, config.baseURL)
	if err != nil {
		fmt.Println("error crawling site:", err)
		return
	}

	for _, nextURL := range foundURLs {
		config.wg.Add(1)
		go config.crawlPage(nextURL)
	}
}
