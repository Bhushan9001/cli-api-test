package cmd

import (
	"github.com/Bhushan9001/cli-api-test/helper"
	"github.com/spf13/cobra"
)

var postURL string
var postBody string
var postHeaders []string

var postCmd = &cobra.Command{

	Use:   "post",
	Short: "Send a POST Request",
	Run: func(cmd *cobra.Command, args []string) {

		if postURL == "" {
			cmd.Usage()
			return
		}
		helper.SendRequest("POST", postURL, postHeaders, postBody, prettyPrint)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.Flags().StringVar(&postURL, "url", "", "Request URL (required)")
	postCmd.Flags().StringVar(&postBody, "body", "", "Request body")
	postCmd.Flags().StringArrayVar(&postHeaders, "header", []string{}, "Request Headers")
	postCmd.Flags().BoolVar(&prettyPrint, "pretty", false, "Formatt json output")
}
