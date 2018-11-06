package questionratio

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/royvandewater/slack-stats/slack"
)

// Ratio represents the ratio of questions to statements
type Ratio struct {
	Numerator   int
	Denominator int
}

func (r *Ratio) String() string {
	return fmt.Sprintf("%v / %v", r.Numerator, r.Denominator)
}

func isQuestion(text string) bool {
	return strings.Contains(text, "?")
}

// FindRatio retrieves the ratio of questions to statements for a user in
// a slack channel
func FindRatio(token, channel, user string, daysAgo int) (*Ratio, error) {
	messages, err := slack.GetUserMessages(token, channel, user, daysAgo)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to getUserMessages")
	}

	numQuestions := 0
	for _, message := range messages {
		if isQuestion(message.Text) {
			numQuestions++
		}
	}

	return &Ratio{numQuestions, len(messages)}, nil
}
