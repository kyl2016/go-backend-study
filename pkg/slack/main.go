package main

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

const (
	webhookUrl = "https://hooks.slack.com/services/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"
)

func main() {
	resp, err := http.Get("url string")

	err := slack.PostWebhook(webhookUrl, &slack.WebhookMessage{
		Channel: "",
		Text:    "*TEST*",
		// Attachments: []slack.Attachment{resultAttachment, confirmAttachment},
	})
	fmt.Println("err:", err)

	// alarm := "TEST"
	// resp, _ := http.Post(webhookUrl, "application/json", strings.NewReader(fmt.Sprintf(`{"text":"%s @channel"}`, alarm)))
	// data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data))
}
