package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"
var cfgPath string

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "This is the first command",
	Long: `A longer description 
	for the first command`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the root command")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(srvCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
