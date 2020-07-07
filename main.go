package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"math"
	"os"
	"time"
)

func main() {
	// My last day is Wednesday July 15, 2020 so that is the end date
	var enddate, _ = time.Parse(time.RFC822, "15 Jul 20 10:00 UTC")
	var slacktoken = os.Getenv("SLACKTOKEN")
	var slackchannel = os.Getenv("SLACKCHANNEL")

	api := slack.New(slacktoken)
	var numdays = calculateWorkingDays(time.Now(), enddate)
	if numdays == 0 {
		var msg = "This is james' last day!!"
		api.PostMessage(slackchannel, slack.MsgOptionText(msg, false))
	} else {
		var msg = fmt.Sprintf("%d working days left before james' last day\n", numdays)
		api.PostMessage(slackchannel, slack.MsgOptionText(msg, false))
	}
}

func calculateWorkingDays(startTime time.Time, endTime time.Time) int {
	// Reduce dates to previous Mondays
	startOffset := weekday(startTime)
	startTime = startTime.AddDate(0, 0, -startOffset)
	endOffset := weekday(endTime)
	endTime = endTime.AddDate(0, 0, -endOffset)

	// Calculate weeks and days
	dif := endTime.Sub(startTime)
	weeks := int(math.Round((dif.Hours() / 24) / 7))
	days := -min(startOffset, 5) + min(endOffset, 5)

	// Calculate total days
	return weeks*5 + days
}

func weekday(d time.Time) int {
	wd := d.Weekday()
	if wd == time.Sunday {
		return 6
	}
	return int(wd) - 1
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
