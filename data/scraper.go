package data

import "fmt"

func main() {
	data, err := scrapers.ScrapeData()

	if err != nil {
		fmt.Println("Error while scraping data:", err)
		return
	}

	// Do something with the data
	for i, item := range data {
		fmt.Println(i+1, ":", item)
	}
	}
}
