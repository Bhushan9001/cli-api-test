package cmd

import (
	"github.com/Bhushan9001/cli-api-test/helper"
	"github.com/spf13/cobra"
)

var putURL string
var putBody string
var putHeaders []string

var putCmd = &cobra.Command{

	Use:   "put",
	Short: "Send a POST Request",
	Run: func(cmd *cobra.Command, args []string) {

		if putURL == "" {
			cmd.Usage()
			return
		}
		helper.SendRequest("PUT", putURL, putHeaders, putBody, prettyPrint)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.Flags().StringVar(&putURL, "url", "", "Request URL (required)")
	putCmd.Flags().StringVar(&putBody, "body", "", "Request body")
	putCmd.Flags().StringArrayVar(&putHeaders, "header", []string{}, "Request headers")
	putCmd.Flags().BoolVar(&prettyPrint, "pretty", false, "Formatt json output")
}
