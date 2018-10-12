package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/royvandewater/slack-stats/ratio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ratioCmd represents the ratio command
var ratioCmd = &cobra.Command{
	Use:   "ratio",
	Short: "Return a ratio of questions to statements.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		channel := viper.GetString("channel")
		user := viper.GetString("user")
		daysAgo := viper.GetInt("days-ago")

		r, err := ratio.FindRatio(token, channel, user, daysAgo)
		if err != nil {
			log.Fatalln("Error occured: ", err.Error())
		}

		t := time.Now().Add(-24 * time.Hour * time.Duration(daysAgo))
		fmt.Printf("%d-%02d-%02d\t%d\t%d\n", t.Year(), t.Month(), t.Day(), r.Numerator, r.Denominator)
	},
}

func init() {
	rootCmd.AddCommand(ratioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	ratioCmd.PersistentFlags().StringP("channel", "c", "", "Slack channel to analyze")
	ratioCmd.PersistentFlags().StringP("user", "u", "", "Slack user to analyze")
	ratioCmd.PersistentFlags().IntP("days-ago", "d", 1, "Days ago the day was to aggregate. 0 aggregates today, 1 yesterday, etc.")

	viper.BindPFlag("channel", ratioCmd.PersistentFlags().Lookup("channel"))
	viper.BindPFlag("user", ratioCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("days-ago", ratioCmd.PersistentFlags().Lookup("days-ago"))
}
