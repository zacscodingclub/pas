package main

import (
	"fmt"
	"regexp"
	"strings"
)

type result struct {
	id           string
	title        string
	state        string
	timeLeft     string
	currentPrice int64
}

func (r *result) toString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{ id: %s,", r.id))
	sb.WriteString(fmt.Sprintf(" title: %s,", r.title))
	sb.WriteString(fmt.Sprintf(" state: %s,", r.state))
	sb.WriteString(fmt.Sprintf(" timeLeft: %s,", r.timeLeft))
	sb.WriteString(fmt.Sprintf(" currentPrice: %d }", r.currentPrice))
	return sb.String()
}

// price, title, time left
func (r *result) toSlack() string {
	var sb strings.Builder
	sb.WriteString(replaceDoubleQuote(fmt.Sprintf("$%d (%s)", r.currentPrice, r.state)))
	sb.WriteString(replaceDoubleQuote(fmt.Sprintf(" <%s%s%s|%s>", c.BaseURL, c.ItemPath, r.id, r.title)))
	sb.WriteString(replaceDoubleQuote(fmt.Sprintf(" %s\n", r.timeLeft)))
	return sb.String()
}

func replaceDoubleQuote(s string) string {
	re := regexp.MustCompile("\"")
	return re.ReplaceAllString(s, "in. ")
}
