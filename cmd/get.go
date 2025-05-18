package cmd

import (
	"github.com/Bhushan9001/cli-api-test/helper"
	"github.com/spf13/cobra"
)

var getURL string
var getHeaders []string
var prettyPrint bool

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Sends a get request and prints Output",
	Run: func(cmd *cobra.Command, args []string) {
		if getURL == "" {
			cmd.Usage()
			return
		}
		helper.SendRequest("GET", getURL, getHeaders, "", prettyPrint)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVar(&getURL, "url", "", "Request URL (Reuired)")
	getCmd.Flags().StringArrayVar(&getHeaders, "header", []string{}, "Request headers")
	getCmd.Flags().BoolVar(&prettyPrint, "pretty", false, "Print the JSON output in proper format")
}
