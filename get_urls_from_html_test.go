package main

import (
	"net/url"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://camescasse.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Camescasse.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Camescasse.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://camescasse.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "empty path",
			inputURL: "https://camescasse.dev",
			inputBody: `
			<html>
				<body>
					<a href="/">
						<span>Camescasse.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://camescasse.dev/"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inputURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v FAIL: input error: %v", tc.name, err)
			}

			actual, err := getURLsFromHTML(tc.inputBody, inputURL)
			if err != nil {
				t.Errorf("Test %v = '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Test %v = '%s' FAIL: expected []URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				}
			}
		})
	}
}
