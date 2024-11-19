package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (config *config) crawlPage(rawCurrentURL string) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing current url:", err)
		return
	}
	if config.baseURL.Host != currentURL.Host {
		return
	}

	currentNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizing current url:", err)
		return
	}

	if _, ok := config.pages[currentNormalized]; ok {
		config.pages[currentNormalized]++
		return
	} else {
		config.pages[currentNormalized] = 1
	}

	res, err := http.Get(rawCurrentURL)
	if err != nil {
		fmt.Println("error getting http body:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		fmt.Printf("error fetching %s: status code: %d\n", rawCurrentURL, res.StatusCode)
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading data:", err)
		return
	}

	htmlContent := string(data)

	fmt.Println("crawling:", rawCurrentURL)
	result, err := getURLsFromHTML(htmlContent, rawCurrentURL)
	if err != nil {
		fmt.Println("error crawling site:", err)
		return
	}

	for _, page := range result {
		config.crawlPage(page)
	}
}
