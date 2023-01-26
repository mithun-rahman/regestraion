package config

import (
	"RegLog/app/response"
	"github.com/spf13/viper"
	"sync"
)

var postgresOnce = sync.Once{}
var postgresConfig *response.Postgres
var srvrInfo *response.ServerInfo

// loadPostgres loads config from path
func loadPostgres(fileName string) error {
	readConfig(fileName)
	viper.AutomaticEnv()

	srvrInfo = &response.ServerInfo{
		ReadTimeout:       viper.GetInt("serverInf.ReadTimeOut"),
		WriteTimeout:      viper.GetInt("serverInf.WriteTimeout"),
		IdleTimeout:       viper.GetInt("serverInf.IdleTimeout"),
		ReadHeaderTimeout: viper.GetInt("serverInf.ReadHeaderTimeout"),
		GracePeriod:       viper.GetInt("serverInf.GracePeriod"),
	}

	postgresConfig = &response.Postgres{
		User:     viper.GetString("postgres.User"),
		Password: viper.GetString("postgres.Password"),
		Name:     viper.GetString("postgres.Name"),
		DbPort:   viper.GetString("postgres.Port"),
		AppPort:  viper.GetInt("app.Port"),
		Host:     viper.GetString("app.Host"),
	}
	return nil
}

// GetloadPostgres returns postgres config
func GetPostgres(fileName string) (*response.Postgres, *response.ServerInfo) {
	postgresOnce.Do(func() {
		loadPostgres(fileName)
	})
	return postgresConfig, srvrInfo
}
