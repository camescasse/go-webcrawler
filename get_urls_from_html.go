package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse url: %w", err)
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse html: %w", err)
	}

	urls := []string{}
	var traverseNode func(*html.Node)
	traverseNode = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v", anchor.Val, err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNode(child)
		}
	}
	traverseNode(doc)

	return urls, nil
}
