package main

import (
	"fmt"
	"strings"
)

type result struct {
	id           string
	title        string
	state        string
	timeLeft     string
	bids         int64
	currentPrice int64
}

func (r *result) toString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{ id: %s,", r.id))
	sb.WriteString(fmt.Sprintf(" title: %s,", r.title))
	sb.WriteString(fmt.Sprintf(" state: %s,", r.state))
	sb.WriteString(fmt.Sprintf(" timeLeft: %s,", r.timeLeft))
	sb.WriteString(fmt.Sprintf(" bids: %d,", r.bids))
	sb.WriteString(fmt.Sprintf(" currentPrice: %d }", r.currentPrice))
	return sb.String()
}

// price, title, time left (bids)
func (r *result) toSlack() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("$%d,", r.currentPrice))
	sb.WriteString(fmt.Sprintf(" <%s%s%s|%s>,", c.BaseURL, c.ItemPath, r.id, r.title))
	sb.WriteString(fmt.Sprintf(" %s", r.timeLeft))
	sb.WriteString(fmt.Sprintf(" (%d bids)\n", r.bids))
	return sb.String()
}
