package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/slack-go/slack"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Endpoints []struct {
		Url     string  `yaml:"url"`
		Timeout float64 `yaml:"timeout"`
	} `yaml:"endpoints"`
}

func getConf(c *conf) {
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func SendSlackMessage(message string, failReason string) {
	slackWebHook := os.Getenv("SLACKWEBHOOK")
	attachment := slack.Attachment{
		Color:  "danger",
		Text:   message,
		Footer: failReason,
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(slackWebHook, &msg)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var configs conf
	getConf(&configs)
	for _, config := range configs.Endpoints {
		client := http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		}
		resp, err := client.Get(config.Url)
		if err != nil {
			var urlError *url.Error
			if errors.As(err, &urlError) {
				if urlError.Timeout() {
					SendSlackMessage(config.Url+" failed", "timeout")
					continue
				}
			}
			log.Println(err)
			SendSlackMessage(config.Url+" failed", err.Error())
		} else {
			defer resp.Body.Close()
			log.Println(config.Url, resp.StatusCode)
			if resp.StatusCode != 200 {
				SendSlackMessage(config.Url+" failed", fmt.Sprintf("StatusCode: %d", resp.StatusCode))
			}
		}
	}
}
