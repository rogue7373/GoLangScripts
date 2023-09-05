package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Define command-line flags for the website URL and output file
	var (
		url        string
		outputFile string
	)

	flag.StringVar(&url, "url", "", "Website URL to scrape")
	flag.StringVar(&outputFile, "output", "output.txt", "Output file to save results")
	flag.Parse()

	// Check if the URL flag is empty
	if url == "" {
		fmt.Println("Please provide a website URL using the -url flag.")
		os.Exit(1)
	}

	// Create a new Colly collector
	c := colly.NewCollector()

	// Create a file to write the results
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		// Extract the title from the HTML element
		title := e.Text
		// Write the title to the output file
		_, err := file.WriteString(title + "\n")
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
		}
		// Print or process the title as needed
		fmt.Println(title)
	})

	// Start scraping by visiting the specified URL
	c.Visit(url)
}
