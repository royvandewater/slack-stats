package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Message is a slack message
type Message struct {
	Text string `json:"text"`
	User string `json:"user"`
}

func startTime(now time.Time, daysAgo int) time.Time {
	t := now.Add(-24 * time.Hour * time.Duration(daysAgo))
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func endTime(now time.Time, daysAgo int) time.Time {
	return startTime(now, daysAgo).Add(24 * time.Hour)
}

func getMessages(token, channel string, daysAgo int) ([]Message, error) {
	slackURL, err := url.Parse("https://slack.com/api/channels.history")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse URL")
	}

	now := time.Now()
	query := slackURL.Query()
	query.Add("channel", channel)
	query.Add("count", "1000")
	query.Add("oldest", fmt.Sprintf("%v.000100", startTime(now, daysAgo).Unix()))
	query.Add("latest", fmt.Sprintf("%v.000100", endTime(now, daysAgo).Unix()))

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

	parsed := struct {
		Messages []Message `json:"messages"`
		OK       bool      `json:"ok"`
		Error    string    `json:"error"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode response")
	}

	if !parsed.OK {
		return nil, fmt.Errorf("Received error from slack: %v", parsed.Error)
	}

	return parsed.Messages, nil
}

// GetUserMessages retrieves messages from a particular user in a channel
func GetUserMessages(token, channel, user string, daysAgo int) ([]Message, error) {
	allMessages, err := getMessages(token, channel, daysAgo)
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
