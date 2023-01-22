package cmd

import (
	"RegLog/server"
	"fmt"
	"github.com/spf13/cobra"
)

var port string

func init() {
	//	startCmd.Flags().StringVarP(&port, "port", "p", "3333", "start the server")
	startCmd.Flags().StringVarP(&port, "port", "p", "3333", "This flag sets the serverPort of the server")
}

var startCmd = &cobra.Command{
	Use:     "serve",
	Short:   "start the server",
	Aliases: []string{"begin, stat"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(port)
		server.Server(port)
	},
}
