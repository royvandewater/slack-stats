package ratio

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/royvandewater/slack-stats/slack"
)

func isQuestion(text string) bool {
	return strings.Contains(text, "?")
}

// FindRatio retrieves the ratio of questions to statments for a user in
// a slack channel
func FindRatio(token, channel, user string) (float64, error) {
	messages, err := slack.GetUserMessages(token, channel, user)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to getUserMessages")
	}

	numQuestions := float64(0)
	for _, message := range messages {
		if isQuestion(message.Text) {
			numQuestions++
		}
	}

	return numQuestions / float64(len(messages)), nil
}
