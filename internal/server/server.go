package server

import (
	"authJWT/configs"
	handler "authJWT/internal/endpoint/handlers"
	"authJWT/internal/middleware"
	"authJWT/internal/router"
	"net/http"
)

func Start() {
	cfg := configs.NewConfig()

	authHandler := handler.NewAuthHandler(cfg)
	userHandler := handler.NewUserHandler(cfg)

	router := router.NewRouter()

	router.AddRoute("/login", authHandler.Login, nil)
	router.AddRoute("/signup", authHandler.Register, nil)
	router.AddRoute("/profile", userHandler.GetProfile, middleware.CheckAccessTokenValidity)
	router.AddRoute("/refresh", userHandler.Refresh, nil)

	http.ListenAndServe(":8080", router)
}
