package postgres

import (
	"RegLog/app/response"
	"RegLog/model"
	"errors"
	"fmt"
	"github.com/alexedwards/argon2id"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Postgres struct {
	DB *gorm.DB
}

func PostNew(config response.Postgres) (*Postgres, error) {

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

func (p *Postgres) Insert(user *response.Users, w http.ResponseWriter) error {
	resp := response.Response{500, "user exits", nil}

	regErr := p.DB.Where("user_name = ?", user.User_name)

	if regErr == nil {
		resp := model.Response{500, "user already exits", nil}
		resp.ServeJSON(w)
		return errors.New("user already exits")
	}

	userPassword := user.Password
	hash, err := argon2id.CreateHash(userPassword, argon2id.DefaultParams)

	if err != nil {
		resp := model.Response{500, "problem with hash function", nil}
		resp.ServeJSON(w)
		return err
	}

	user.Password = hash

	if er := p.DB.Create(&user).Error; er != nil {
		resp.ServeJSON(w)
		return er
	}
	return nil
}

func (p *Postgres) Get(user *response.Users, w http.ResponseWriter) (response.Users, error) {
	resp := response.Response{500, "server problem", nil}

	var fromDB, empUser response.Users
	password := user.Password

	logErr := p.DB.Where("user_name = ?", user.User_name).First(&fromDB).Error
	if logErr != nil {
		resp := model.Response{404, "user not found", nil}
		resp.ServeJSON(w)
		return empUser, logErr
	}
	ok, err := argon2id.ComparePasswordAndHash(password, fromDB.Password)
	if err != nil {
		resp.ServeJSON(w)
		return empUser, err
	}
	if ok {
		return fromDB, nil
	} else {
		return empUser, nil
	}
}
