package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Endpoints []struct {
		Url     string `yaml:"url"`
		Timeout float64 `yaml:"timeout"`
	} `yaml:"endpoints"`
}

func getConf(c *conf) {
    yamlFile, err := ioutil.ReadFile("conf.yaml")
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
		Color:         "danger",
		Text:          message,
		Footer:        failReason,
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
	for _, config := range configs.Endpoints{
		start := time.Now()
		resp, err := http.Get(config.Url)
		elapsed := time.Since(start).Seconds()
		defer resp.Body.Close()
		if err != nil {
			log.Println(err)
			SendSlackMessage(config.Url + " failed", "Didn't get any response")
		} else {
			log.Println(config.Url, resp.StatusCode)
			if resp.StatusCode != 200 {
				SendSlackMessage(config.Url + " failed", "StatusCode:" + strconv.Itoa(resp.StatusCode))
			}
			if elapsed > config.Timeout {
				SendSlackMessage(config.Url + " failed", "Reply toke " + strconv.FormatFloat(elapsed,'E', -1, 32))
			}
			log.Println(elapsed)
		}
	}
}
