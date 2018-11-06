package userratio

import (
	"fmt"

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

// FindRatio retrieves the ratio of questions to statements for a user in
// a slack channel
func FindRatio(token, channel, user string, daysAgo int) (*Ratio, error) {
	messages, err := slack.GetMessages(token, channel, daysAgo)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to getMessages")
	}

	numUserMessages := 0
	for _, message := range messages {
		if message.User == user {
			numUserMessages++
		}
	}

	return &Ratio{numUserMessages, len(messages)}, nil
}
