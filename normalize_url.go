package main

import (
	"net/url"
	"strings"
)

func normalizeURL(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	normalized := u.Host + u.Path
	normalized = strings.ToLower(normalized)
	normalized = strings.TrimSuffix(normalized, "/")

	return normalized, nil
}
