package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	// Expect two URLs in argv, a URL and filename of a vector map.
	num_args := len(os.Args)
	if num_args < 3 {
		fmt.Println("Too few arguments. (ex: wordmapper vector_map https://example.com)")
		return
	}
	if num_args > 3 {
		fmt.Println("Too many arguments. (ex: wordmapper vector_map https://example.com)")
		return
	}


	// Crawl URL
	// Fetch HTML and parse into single doc string
	// Doc to cleaned words
	// Cleaned words to map with occurrences as values
	// Occurrence map to importance map, calculated with vector values
	// Visualize importance map.

	// Start web server
	// Serve app that accepts argv via form entry
	// Render visualized data on page
}
