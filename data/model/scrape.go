package model

import (
	"fmt"
	"log"
	"strings"

	"github.com/HousePricePrediction/data/model"
	"github.com/PuerkitoBio/goquery"
)

// ScrapeProperties scrapes properties data from the given URLs
func ScrapeProperties(urls []string) []model.Property {
	var properties []model.Property

	for _, url := range urls {
		// Request the HTML page.
		res, err := Request(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Scrape data from the first website
		if strings.Contains(url, "first-website") {
			doc.Find(".property-card").Each(func(i int, s *goquery.Selection) {
				price := s.Find(".property-card__price").Text()
				price = strings.TrimSpace(price)

				location := s.Find(".property-card__location").Text()
				location = strings.TrimSpace(location)

				property := model.Property{
					Price:    price,
					Location: location,
				}
				properties = append(properties, property)
			})
		}

		// Scrape data from the second website
		if strings.Contains(url, "second-website") {
			doc.Find(".result-search__item__description").Each(func(i int, s *goquery.Selection) {
				price := s.Find(".result-search__item__price").Text()
				price = strings.TrimSpace(price)

				location := s.Find(".result-search__item__title").Text()
				location = strings.TrimSpace(location)

				property := model.Property{
					Price:    price,
					Location: location,
				}
				properties = append(properties, property)
			})
		}
	}

	return properties
}

func main() {
	urls := []string{
		"https://www.first-website.com/property",
		"https://www.second-website.com/property",
	}

	properties := ScrapeProperties(urls)
	for _, property := range properties {
		fmt.Printf("Price: %s, Location: %s\n", property.Price, property.Location)
	}
}
