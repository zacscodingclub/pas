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

func (r *result) toSlack() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{ id: %s,", r.id))
	sb.WriteString(fmt.Sprintf(" title: %s,", r.title))
	sb.WriteString(fmt.Sprintf(" state: %s,", r.state))
	sb.WriteString(fmt.Sprintf(" timeLeft: %s,", r.timeLeft))
	sb.WriteString(fmt.Sprintf(" bids: %d,", r.bids))
	sb.WriteString(fmt.Sprintf(" currentPrice: %d }", r.currentPrice))
	return sb.String()
}
