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

	//r.Post("/reg", middle.Limit(con.Register))
	//r.Post("/log", middle.Limit(con.Login))
	//r.Get("/user", middle.Limit(con.GetUser))
	//r.Post("/refresh", middle.Limit(con.RefreshToken))

	r.Post("/reg", con.Register)
	r.Post("/log", con.Login)
	r.Get("/user", con.GetUser)
	r.Post("/refresh", con.RefreshToken)

	port = ":" + port
	http.ListenAndServe(port, r)

}
