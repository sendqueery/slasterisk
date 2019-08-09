package slasterisk

import (
	"encoding/json"
	"log"
	"net/http"
)

type ChannelData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsChannel bool   `json:"is_channel"`
	IsMember  bool   `json:"is_member"`
}

type SlackAPIResponse struct {
	Channels []ChannelData `json:"channels"`
}

func joinURL(baseURL string, token string, args ...string) (URL string) {
	URL = (baseURL + "token=" + token)
	if len(args) == 0 {
		return URL
	}

	for _, v := range args[1:] {
		URL += ("&" + v)
	}
	return URL
}

func GetSlackChannels(token string) (channels *SlackAPIResponse) {
	// URL for the Slack conversation API
	baseURL := "https://slack.com/api/conversations.list?"
	// Build our API call URL
	URL := joinURL(baseURL, token)

	// Make our API call
	resp := CallSlackAPI(URL)
	defer resp.Body.Close()

	// Set up our JSON decoding and response body struct
	body := json.NewDecoder(resp.Body)
	channelInfo := new(SlackAPIResponse)

	for body.More() {
		if err := body.Decode(&channelInfo); err != nil {
			log.Fatal("Invalid JSON detected.")
		}
	}
	return channelInfo
}

func CallSlackAPI(URL string) (resp *http.Response) {
	resp = nil
	client := &http.Client{}
	if resp, err := client.Get(URL); err != nil || resp.StatusCode != 200 {
		log.Fatalf("API call failed or returned non-200 status code\nException: %v", err)
		return nil
	} else {
		return resp
	}
}

func ValidateSlackChannel(channelName string, channels *SlackAPIResponse) (channelID string) {
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
