package main

import (
	"fmt"
	"os"

	"github.com/benskia/WordMapper/internal/weights"
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

	mymap := weights.BuildWeightMap(os.Args[1])
	fmt.Print(mymap)

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
