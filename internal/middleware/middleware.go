package middleware

import (
	"authJWT/configs"
	"authJWT/internal/service"
	"context"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

type contextKey string

const (
	contextClaimKey contextKey = "claims"
)

func CheckAccessTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ats := configs.NewConfig().AccessTokenSecret
		claims, err := service.GetClaims(w, r, ats)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(context.Background(), contextClaimKey, claims)
		rctx := r.WithContext(ctx)

		if claims != nil {
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, rctx)
			return
		}
	})
}
