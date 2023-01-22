package cmd

import (
	"RegLog/config"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"

	infraPostgres "RegLog/infra/postgres"
)

// srvCmd is the serve sub command to start the api server
var srvCmd = &cobra.Command{
	Use:   "serv",
	Short: "serve serves the api server",
	RunE:  serve,
}

func init() {
	srvCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "config file path")
}

func serve(cmd *cobra.Command, args []string) error {
	cfgPostgres := config.GetPostgres(cfgPath)
	db, err := infraPostgres.PostNew(*cfgPostgres)
	if err != nil {
		fmt.Println("problem database connection")
		return err
	}
	fmt.Println(db)
	http.ListenAndServe(":3333", nil)
	return nil
}
