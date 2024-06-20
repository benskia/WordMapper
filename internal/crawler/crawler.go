package crawler

import (
	"fmt"
	"net/url"
)

func CrawlPage(baseURL, currentURL string, pages []string) []string {
	baseURLObj, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing baseURL: %v", baseURL)
		return pages
	}

	currentURLObj, err := url.Parse(currentURL)
	if err != nil {
		fmt.Printf("Error parsing currentURL: %v", baseURL)
		return pages
	}

	// We don't want to leave the targeted host.
	if baseURLObj.Hostname() != currentURLObj.Hostname() {
		return pages
	}

	return pages
}
