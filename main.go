package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	ask "github.com/sendqueery/slasterisk/asterisk"
	// slack "github.com/sendqueery/slasterisk/slack"
)

type Config struct {
	SlackInfo struct {
		Token       string `json:"token"`
		ChannelName string `json:"channel_name"`
	}
	AsteriskInfo struct {
		VMDir     string `json:"vm_dir"`
		Extension string `json:"extension"`
	}
}

func main() {

	// Test stuff for JSON parsing
	configData := getJSONConfig()
	test, _ := json.MarshalIndent(configData, "", "\t")
	fmt.Printf("%v\n", string(test))

	api := slack.New(configData.SlackInfo.Token)

	cp := slack.GetConversationsParameters{
		ExcludeArchived: "true",
		Types:           []string{"public_channel"},
	}

	var channelID string
	for channels, nextCursor, err := api.GetConversations(&cp); channelID == ""; {

		if err != nil {
			fmt.Println(err)
		} else {
			for _, v := range channels {
				if v.Name == configData.SlackInfo.ChannelName && v.IsMember == true {
					channelID = v.ID
				}
			}
			if nextCursor != "" {
				cp.Cursor = nextCursor
			} else if nextCursor == "" && channelID == "" {
				log.Fatalf(`Channel "%v" is not a valid Slack channel, or the API token has not been authorized to access it.`, configData.SlackInfo.ChannelName)
			}
		}
	}

	//channelList := slack.GetSlackChannels(configData.SlackInfo.Token)
	//channelID := slack.ValidateSlackChannel(configData.SlackInfo.ChannelName, channelList)
	fmt.Println(channelID)

	// Test stuff for parsing a VM info file
	info := ask.ParseVMInfo(`C:\Users\amaheu\go\src\github.com\sendqueery\slasterisk\test_data\msg0003.txt`)
	fmt.Println(info)
}

func getJSONConfig() (configData *Config) {
	// Initialize our empty struct
	data := new(Config)

	// Make sure our config file exists
	if file, err := os.Open(`C:\Users\amaheu\go\src\github.com\sendqueery\slasterisk\config\config.json`); err != nil {
		log.Fatalln(err)
	} else {
		// Don't close our reader until we're done here
		defer file.Close()

		// Use a buffered reader and decode the JSON
		b := bufio.NewReader(file)
		rawcontent := json.NewDecoder(b)

		// While we have more data, keep parsing
		for rawcontent.More() {
			if err := rawcontent.Decode(&data); err != nil {
				log.Fatal("Invalid JSON detected.")
			}
		}
	}
	// Return our populated struct
	return data
}
