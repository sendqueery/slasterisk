package slasterisk

import (
	"log"
	"net/http"
)

type slackChannelAPIResponse struct {
	Channels []ChannelData `json:"channels"`
}

type ChannelData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsChannel bool   `json:"is_channel"`
	IsMember  bool   `json:"is_member"`
}

type slackUserAPIResponse struct {
	Users []UserData `json:"members"`
}

type UserData struct {
	ID      string            `json:"id"`
	Profile []UserProfileData `json:"profile"`
}

type UserProfileData struct {
	Phone string `json:"phone"`
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
