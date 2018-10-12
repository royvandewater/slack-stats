package ratio

import (
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"github.com/royvandewater/slack-stats/slack"
)

func isQuestion(text string) bool {
	return strings.Contains(text, "?")
}

// FindRatio retrieves the ratio of questions to statments for a user in
// a slack channel
func FindRatio(token, channel, user string) (*big.Rat, error) {
	messages, err := slack.GetUserMessages(token, channel, user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to getUserMessages")
	}

	numQuestions := int64(0)
	for _, message := range messages {
		if isQuestion(message.Text) {
			numQuestions++
		}
	}

	return big.NewRat(numQuestions, int64(len(messages))), nil
}
