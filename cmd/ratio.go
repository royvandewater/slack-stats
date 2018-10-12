package cmd

import (
	"fmt"
	"log"

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

		r, err := ratio.FindRatio(token, channel, user)
		if err != nil {
			log.Fatalln("Error occured: ", err.Error())
		}

		fmt.Println(r.String())
	},
}

func init() {
	rootCmd.AddCommand(ratioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	ratioCmd.PersistentFlags().StringP("channel", "c", "", "Slack channel to analyze")
	ratioCmd.PersistentFlags().StringP("user", "u", "", "Slack user to analyze")

	viper.BindPFlag("channel", ratioCmd.PersistentFlags().Lookup("channel"))
	viper.BindPFlag("user", ratioCmd.PersistentFlags().Lookup("user"))
}
