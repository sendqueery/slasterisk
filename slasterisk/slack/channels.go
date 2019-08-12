package slasterisk

import (
	"encoding/json"
	"log"
)

func GetSlackChannels(token string) (channels *slackChannelAPIResponse) {
	// URL for the Slack conversation API
	baseURL := "https://slack.com/api/conversations.list?"
	// Build our API call URL
	URL := joinURL(baseURL, token)

	// Make our API call
	resp := CallSlackAPI(URL)
	defer resp.Body.Close()

	// Set up our JSON decoding and response body struct
	body := json.NewDecoder(resp.Body)
	channelInfo := new(slackChannelAPIResponse)

	for body.More() {
		if err := body.Decode(&channelInfo); err != nil {
			log.Fatal("Invalid JSON detected.")
		}
	}
	return channelInfo
}

func ValidateSlackChannel(channelName string, channels *slackChannelAPIResponse) (channelID string) {
	channelID = ""
	for _, v := range channels.Channels {
		if v.Name == channelName && v.IsMember && v.IsChannel {
			channelID = v.ID
		}
	}
	if channelID == "" {
		log.Fatalf("Cannot find channel, or app is not authorized to post to channel: %v", channelName)
	}
	return channelID
}
