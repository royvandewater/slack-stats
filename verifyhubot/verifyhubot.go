package verifyhubot

import (
	"fmt"
	"strings"

	"github.com/royvandewater/slack-stats/slack"
)

// Command represents a single attempt to enable a trial
type Command struct {
	Org string
	Success bool
}

func isHubot(user string) bool {
	return user == "U0JJLJEBC"
}

func isSuccess(text string) bool {
	return strings.Contains(text, ":trial {")
}

// VerifyHubot verifies that a hubot command
// actually enabled a trial plan
func VerifyHubot(token, channel string) ([]*Command, error) {
	messages, err := slack.GetMessages(token, channel, 100)

	if err != nil {
		return nil, err
	}

	commands := make([]*Command, 0)

	for _, message := range messages {
		if !isHubot(message.User) {
			continue
		}

		if len(message.Files) < 1 {
			continue
		}

		file := message.Files[0]
		contents, err := slack.GetFileContents(token, file.URLPrivateDownload)
		if err != nil {
			fmt.Println("error", err.Error())
			continue
		}

		command := &Command{
			Org: file.Title,
			Success: isSuccess(contents),
		}
		commands = append(commands, command)
	}

	return commands, nil
}