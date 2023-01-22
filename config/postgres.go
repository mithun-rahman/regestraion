package config

import (
	"github.com/spf13/viper"
	"sync"
)

// Postgres holds postgres config
type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	DbPort   string `yaml:"dbPort"`
	AppPort  string `yaml:"appPort"`
}

var postgresOnce = sync.Once{}
var postgresConfig *Postgres

// loadPostgres loads config from path
func loadPostgres(fileName string) error {
	readConfig(fileName)
	viper.AutomaticEnv()

	postgresConfig = &Postgres{
		User:     viper.GetString("postgres.User"),
		Password: viper.GetString("postgres.Password"),
		Name:     viper.GetString("postgres.Name"),
		DbPort:   viper.GetString("postgres.Port"),
		AppPort:  viper.GetString("app.Port"),
	}
	return nil
}

// GetloadPostgres returns postgres config
func GetPostgres(fileName string) *Postgres {
	postgresOnce.Do(func() {
		loadPostgres(fileName)
	})
	return postgresConfig
}
