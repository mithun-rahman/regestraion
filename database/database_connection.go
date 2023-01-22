package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func IntializeDB() *gorm.DB {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	user := viper.GetString("postgres.User")
	password := viper.GetString("postgres.Password")
	dbname := viper.GetString("postgres.Name")
	port := viper.GetString("postgres.Port")

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("user=%s password=%s host=127.0.0.1 port=%v dbname=%s sslmode=disable", user, password, port, dbname),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("not connected")
		log.Panic("connection problem")
	}

	fmt.Println("connected")

	return conn
}
