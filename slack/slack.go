package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"

	"github.com/pkg/errors"
)

// File is a file attached to a message
type File struct {
	URLPrivateDownload string `json:"url_private_download"`
	Title string `json:"title"`
}

// Message is a slack message
type Message struct {
	Text string `json:"text"`
	User string `json:"user"`
	Files []File `json:"files"`
}

// GetMessages retrieves messages in a channel
func GetMessages(token, channel string, daysAgo int) ([]Message, error) {
	slackURL, err := url.Parse("https://slack.com/api/conversations.history")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse URL")
	}

	// now := time.Now()
	query := slackURL.Query()
	query.Add("channel", channel)
	query.Add("count", "1000")

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
	allMessages, err := GetMessages(token, channel, daysAgo)
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

func GetFileContents(token, urlPrivateDownload string) (string, error) {
	slackURL, err := url.Parse(urlPrivateDownload)
	if err != nil {
		return "", errors.Wrap(err, "Failed to parse URL")
	}


	req, err := http.NewRequest("GET", slackURL.String(), nil)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create request")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "Failed to Do request")
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Non 200 status code returned: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read the response body")
	}

	return string(data), nil
}