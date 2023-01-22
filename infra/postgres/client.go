package postgres

import (
	"RegLog/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	DB *gorm.DB
}

func PostNew(config config.Postgres) (*Postgres, error) {

	user := config.User
	password := config.Password
	dbname := config.Name
	port := config.DbPort

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("user=%s password=%s host=127.0.0.1 port=%v dbname=%s sslmode=disable", user, password, port, dbname),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("not connected")
		log.Panic("connection problem")
	}

	post := &Postgres{db}
	fmt.Println("connected")

	return post, nil
}
