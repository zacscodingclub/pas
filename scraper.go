package main

import "log"

type scraper struct {
	Results []result
	config  conf
}

func NewScraper(c conf) *scraper {
	results := make([]result, 50)
	return &scraper{
		Results: results,
		config:  c,
	}
}

func (s *scraper) scrapeIndex() {
	url := s.config.BaseURL + s.config.ScrapePath
	log.Printf("Now Scraping: %s", url)
}
