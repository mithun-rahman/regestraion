package server

import (
	"RegLog/controller"
	"RegLog/database"
	"RegLog/middle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"
	"time"
)

func Server(port string) {

	r := chi.NewRouter()

	con := controller.Controller{}
	con.DB = database.IntializeDB()

	signBytes, _ := ioutil.ReadFile("./app.rsa")
	signKey, _ := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	con.PrivateKey = signKey

	verifyBytes, _ := ioutil.ReadFile("./app.rsa.pub")
	verifyKey, _ := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	con.PublicKey = verifyKey

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middle.Limit)

	r.Post("/reg", con.Register)
	r.Post("/log", con.Login)
	r.Get("/user", con.GetUser)
	r.Post("/refresh", con.RefreshToken)

	port = ":" + port
	srvr := http.Server{
		Addr:    port,
		Handler: r,
		//ErrorLog: logger.DefaultErrLogger,
		//WriteTimeout: cfg.WriteTimeout,
		//ReadTimeout:  cfg.ReadTimeout,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	srvr.ListenAndServe()
	//http.ListenAndServe(port, r)

}
