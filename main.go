package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
)

type conf struct {
	BaseURL      string `json:"baseUrl"`
	ScrapePath   string `json:"scrapePath"`
	ItemPath     string `json:"itemPath"`
	ItemSelector string `json:"itemSelector"`
	WebhookURL   string `json:"webhookUrl"`
}

var c conf

func main() {
	lambda.Start(run)
	// run()
}

func run() {
	c.getConf()

	results := make([]result, 0, 50)

	s := colly.NewCollector()

	s.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s", r.URL)
	})

	skipHeader := true
	s.OnHTML(c.ItemSelector, func(e *colly.HTMLElement) {
		if skipHeader {
			skipHeader = false
			return
		}

		r := buildResultFromElement(e)

		results = append(results, r)
	})

	s.OnScraped(func(r *colly.Response) {
		log.Println("Finished Scraping")
		filteredResults := finishingToday(results)
		msg := buildMessage(filteredResults)
		postToSlack(msg)
	})
	s.Visit(c.BaseURL + c.ScrapePath)
}

func finishingToday(rs []result) []result {
	filtered := make([]result, 0)
	for _, r := range rs {
		if !strings.Contains(r.timeLeft, "day") {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

func buildResultFromElement(e *colly.HTMLElement) result {
	splitPrice := strings.Split(e.ChildText("td:nth-child(6)"), ".")
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	price, err := strconv.ParseInt(reg.ReplaceAllString(splitPrice[0], ""), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return result{
		id:           e.ChildText("td:nth-child(1)"),
		title:        e.ChildText("td:nth-child(2)"),
		state:        e.ChildText("td:nth-child(4)"),
		timeLeft:     e.ChildText("td:nth-child(5)"),
		currentPrice: price,
	}
}

func buildMessage(rs []result) string {
	var sb strings.Builder
	if len(rs) < 1 {
		sb.WriteString(`{"text":"No auctions ending today."}`)
	} else {
		sb.WriteString(fmt.Sprintf(`{"text":"Found %d listings","attachments": [{"text":"`, len(rs)))
		for _, r := range rs {
			sb.WriteString(r.toSlack())
		}
		sb.WriteString(fmt.Sprintf(`"}]}`))
	}
	return sb.String()
}

func postToSlack(r string) {
	req, err := http.NewRequest("POST", c.WebhookURL, strings.NewReader(r))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
}

func (c *conf) getConf() *conf {
	if isLambda := os.Getenv("AWS_EXECUTION_ENV"); len(isLambda) < 1 {
		jsonFile, err := ioutil.ReadFile("./.env.json")
		if err != nil {
			log.Printf("env.json.Get err   #%v ", err)
		}
		err = json.Unmarshal(jsonFile, c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	} else {
		c.BaseURL = os.Getenv("BASE_URL")
		c.ScrapePath = os.Getenv("SCRAPE_PATH")
		c.ItemPath = os.Getenv("ITEM_PATH")
		c.ItemSelector = os.Getenv("ITEM_SELECTOR")
		c.WebhookURL = os.Getenv("WEBHOOK_URL")
	}

	return c
}
