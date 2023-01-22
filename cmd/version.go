package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version of the api server",
	Long:  "This command will print the current version of the api server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
