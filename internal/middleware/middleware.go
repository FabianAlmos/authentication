package middleware

import (
	"authJWT/configs"
	"authJWT/internal/service"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

func CheckAccessTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ats := configs.NewConfig().AccessTokenSecret
		authHeader := r.Header.Get("Authorization")
		tokenString := service.GetTokenFromBearerString(authHeader)
		claims, err := service.ValidateToken(tokenString, ats)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if claims != nil {
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, r)
			return
		}
	})
}
