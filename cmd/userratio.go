package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/royvandewater/slack-stats/userratio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userRatioCommand represents the ratio command
var userRatioCommand = &cobra.Command{
	Use:     "user-ratio",
	Aliases: []string{"u"},
	Short:   "Return a ratio of user messages to all messages in a channel.",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		channel := viper.GetString("channel")
		user := viper.GetString("user")
		daysAgo := viper.GetInt("days-ago")

		r, err := userratio.FindRatio(token, channel, user, daysAgo)
		if err != nil {
			log.Fatalln("Error occured: ", err.Error())
		}

		t := time.Now().Add(-24 * time.Hour * time.Duration(daysAgo))
		fmt.Printf("%d-%02d-%02d\t%d\t%d\n", t.Year(), t.Month(), t.Day(), r.Numerator, r.Denominator)
	},
}

func init() {
	rootCmd.AddCommand(userRatioCommand)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	userRatioCommand.PersistentFlags().StringP("channel", "c", "", "Slack channel to analyze")
	userRatioCommand.PersistentFlags().StringP("user", "u", "", "Slack user to analyze")
	userRatioCommand.PersistentFlags().IntP("days-ago", "d", 1, "Days ago the day was to aggregate. 0 aggregates today, 1 yesterday, etc.")

	viper.BindPFlag("channel", userRatioCommand.PersistentFlags().Lookup("channel"))
	viper.BindPFlag("user", userRatioCommand.PersistentFlags().Lookup("user"))
	viper.BindPFlag("days-ago", userRatioCommand.PersistentFlags().Lookup("days-ago"))
}
