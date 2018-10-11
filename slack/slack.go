package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Message is a slack message
type Message struct {
	Text string `json:"text"`
	User string `json:"user"`
}

func getMessages(token, channel string) ([]Message, error) {
	slackURL, err := url.Parse("https://slack.com/api/channels.history")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse URL")
	}

	query := slackURL.Query()
	query.Add("channel", channel)

	slackURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", slackURL.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create request")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to Do request")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Non 200 status code returned: %v", resp.StatusCode)
	}

	messages := make([]Message, 0)
	err = json.NewDecoder(resp.Body).Decode(&messages)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode response")
	}

	return messages, nil
}

// GetUserMessages retrieves messages from a particular user in a channel
func GetUserMessages(token, channel, user string) ([]Message, error) {
	allMessages, err := getMessages(token, channel)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to getMessages")
	}

	messages := make([]Message, 0)
	for _, message := range allMessages {
		if message.User == user {
			messages = append(messages, message)
		}
	}
	return messages, nil
}
