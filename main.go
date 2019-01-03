package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type conf struct {
	BaseURL      string `json:"baseUrl"`
	ScrapePath   string `json:"scrapePath"`
	ItemPath     string `json:"itemPath"`
	ItemSelector string `json:"itemSelector"`
}

type result struct {
	id           string
	title        string
	state        string
	timeLeft     string
	bids         int64
	currentPrice int64
}

func main() {
	var c conf
	c.getConf()

	results := make([]result, 50)
	skipFirst := true
	s := colly.NewCollector()

	s.OnHTML(c.ItemSelector, func(e *colly.HTMLElement) {
		if skipFirst {
			skipFirst = false
			return
		}
		numBids, err := strconv.ParseInt(e.ChildText("td:nth-child(6)"), 10, 64)
		splitPrice := strings.Split(e.ChildText("td:nth-child(7)"), ".")
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		price, err := strconv.ParseInt(reg.ReplaceAllString(splitPrice[0], ""), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		r := result{
			id:           e.ChildText("td:nth-child(1)"),
			title:        e.ChildText("td:nth-child(2)"),
			state:        e.ChildText("td:nth-child(4)"),
			timeLeft:     e.ChildText("td:nth-child(5)"),
			bids:         numBids,
			currentPrice: price,
		}
		log.Printf("Result: %s", r)

		results = append(results, r)
	})

	s.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s", r.URL)
	})

	s.Visit(c.BaseURL + c.ScrapePath)
}

func (c *conf) getConf() *conf {
	jsonFile, err := ioutil.ReadFile("./.env.json")
	if err != nil {
		log.Printf("env.json.Get err   #%v ", err)
	}

	err = json.Unmarshal(jsonFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println(c.BaseURL)
	return c
}
