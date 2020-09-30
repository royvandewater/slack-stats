package cmd

import (
	"fmt"
	"log"

	"github.com/royvandewater/slack-stats/verifyhubot"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// verifyHubotCommand represents the ratio command
var verifyHubotCommand = &cobra.Command{
	Use:     "verify-hubot",
	Aliases: []string{"q"},
	Short:   "Verifies that the hubot command actually activated the trial",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		channel := viper.GetString("question.channel")

		commands, err := verifyhubot.VerifyHubot(token, channel)
		if err != nil {
			log.Fatalln("Error occured: ", err.Error())
		}

		for _, command := range commands {
			fmt.Printf("%v\t%v\n", command.Success, command.Org)
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyHubotCommand)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	verifyHubotCommand.PersistentFlags().StringP("channel", "c", "", "Slack channel to analyze")

	viper.BindPFlag("question.channel", verifyHubotCommand.PersistentFlags().Lookup("channel"))
}
