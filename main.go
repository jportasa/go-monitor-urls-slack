package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/slack-go/slack"
)

func SendSlackMessage(message string, failReason string) {
	attachment := slack.Attachment{
		Color:         "danger",
		Text:          message,
		Footer:        failReason,

	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook("WEBHOOKURL", &msg)
	if err != nil {
		log.Println(err)
	}	
}

func main() {
	urls := flag.String("u", "", "url")
	responseTimeout := flag.Int("t", 1, "response tiemout")
	flag.Parse()

	for _, url := range urls {
		start := time.Now()
		resp, err := http.Get(url)
		elapsed := time.Since(start).Seconds()
		defer resp.Body.Close()
		if err != nil {
			log.Println(err)
			SendSlackMessage(url + " failed", "Didn't get any response")
		} else {
			log.Println(url, resp.StatusCode)
			if resp.StatusCode != 200 {
				SendSlackMessage(url + " failed", "StatusCode:" + strconv.Itoa(resp.StatusCode))
			}
			if elapsed > responseTimeout {
				SendSlackMessage(url + " failed", "Reply toke " + strconv.FormatFloat(elapsed,'E', -1, 32))
			}
			log.Println(elapsed)
		}
	}
}
