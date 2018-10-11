package cmd

import (
	"fmt"
	"log"

	"github.com/royvandewater/slack-stats/ratio"
	"github.com/spf13/cobra"
)

// ratioCmd represents the ratio command
var ratioCmd = &cobra.Command{
	Use:   "ratio",
	Short: "Return a ratio of questions to statements.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := cmd.Flag("token").Value.String()
		channel := cmd.Flag("channel").Value.String()
		user := cmd.Flag("user").Value.String()

		r, err := ratio.FindRatio(token, channel, user)
		if err != nil {
			log.Fatalln("Error occured: ", err.Error())
		}
		fmt.Println(r)
	},
}

func init() {
	rootCmd.AddCommand(ratioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	ratioCmd.PersistentFlags().StringP("channel", "c", "", "Slack channel to analyze")
	ratioCmd.PersistentFlags().StringP("user", "u", "", "Slack user to analyze")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ratioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ratioCmd.MarkPersistentFlagRequired("channel")
	ratioCmd.MarkPersistentFlagRequired("user")
}
