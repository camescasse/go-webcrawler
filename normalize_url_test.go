package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name:     "remove scheme keep path",
			inputURL: "https://gallery.camescasse.dev/path",
			expected: "gallery.camescasse.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://gallery.camescasse.dev/path/",
			expected: "gallery.camescasse.dev/path",
		},
		{
			name:     "lowercase letters",
			inputURL: "http://GALLERY.camesCasse.Dev/patH/",
			expected: "gallery.camescasse.dev/path",
		},
		{
			name:     "remove scheme without path",
			inputURL: "https://gallery.camescasse.dev/",
			expected: "gallery.camescasse.dev",
		},
		{
			name:     "change nothing",
			inputURL: "gallery.camescasse.dev",
			expected: "gallery.camescasse.dev",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\InvalidURL`,
			expected:      "",
			errorContains: "missing protocol scheme",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v = '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v = '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v = '%s' FAIL: expected error containing %v, got none", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v = '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
