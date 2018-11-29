package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/royvandewater/slack-stats/questionratio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// questionRatioCmd represents the ratio command
var questionRatioCmd = &cobra.Command{
	Use:     "question-ratio",
	Aliases: []string{"q"},
	Short:   "Return a ratio of questions to statements.",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		channel := viper.GetString("question.channel")
		user := viper.GetString("question.user")
		daysAgo := viper.GetInt("question.days-ago")

		r, err := questionratio.FindRatio(token, channel, user, daysAgo)
		if err != nil {
			log.Fatalln("Error occured: ", err.Error())
		}

		t := time.Now().Add(-24 * time.Hour * time.Duration(daysAgo))
		fmt.Printf("%d-%02d-%02d\t%d\t%d\n", t.Year(), t.Month(), t.Day(), r.Numerator, r.Denominator)
	},
}

func init() {
	rootCmd.AddCommand(questionRatioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	questionRatioCmd.PersistentFlags().StringP("channel", "c", "", "Slack channel to analyze")
	questionRatioCmd.PersistentFlags().StringP("user", "u", "", "Slack user to analyze")
	questionRatioCmd.PersistentFlags().IntP("days-ago", "d", 2, "Days ago the day was to aggregate. 0 aggregates today, 1 yesterday, etc.")

	viper.BindPFlag("question.channel", questionRatioCmd.PersistentFlags().Lookup("channel"))
	viper.BindPFlag("question.user", questionRatioCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("question.days-ago", questionRatioCmd.PersistentFlags().Lookup("days-ago"))
}
