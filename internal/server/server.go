package server

import (
	"authJWT/configs"
	"authJWT/internal/endpoint"
	"log"
	"net/http"
)

func Start() {
	cfg := configs.NewConfig()

	authHandler := endpoint.NewAuthHandler(cfg)
	userHandler := endpoint.NewUserHandler(cfg)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", userHandler.GetProfile)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
