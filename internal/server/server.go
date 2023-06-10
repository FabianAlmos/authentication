package server

import (
	"authJWT/configs"
	"authJWT/internal/endpoint"
	"authJWT/internal/middleware"
	"authJWT/internal/router"
	"net/http"
)

func Start() {
	cfg := configs.NewConfig()

	authHandler := endpoint.NewAuthHandler(cfg)
	userHandler := endpoint.NewUserHandler(cfg)

	router := router.NewRouter()

	router.AddRoute("/login", authHandler.Login, nil)
	router.AddRoute("/profile", userHandler.GetProfile, middleware.CheckAccessTokenValidity)
	router.AddRoute("/refresh", userHandler.Refresh, nil)

	http.ListenAndServe(":8080", router)
}
