package cmd

import (
	"github.com/Bhushan9001/cli-api-test/helper"
	"github.com/spf13/cobra"
)

var deleteURL string
var deleteBody string
var deleteHeaders []string

var deleteCmd = &cobra.Command{

	Use:   "delete",
	Short: "Send a POST Request",
	Run: func(cmd *cobra.Command, args []string) {

		if deleteURL == "" {
			cmd.Usage()
			return
		}

		helper.SendRequest("DELETE", deleteURL, deleteHeaders, deleteBody, prettyPrint)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVar(&deleteURL, "url", "", "Request URL (required)")
	deleteCmd.Flags().StringVar(&deleteBody, "body", "", "Request body")
	deleteCmd.Flags().StringArrayVar(&deleteHeaders, "header", []string{}, "Request headers")
	deleteCmd.Flags().BoolVar(&prettyPrint, "pretty", false, "Formatt json output")

}
