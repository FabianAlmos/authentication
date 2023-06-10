package handler

import (
	"authJWT/configs"
	"authJWT/internal/db/repo"
	response "authJWT/internal/endpoint/responses"
	"authJWT/internal/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	cfg *configs.Config
}

func NewUserHandler(cfg *configs.Config) *UserHandler {
	return &UserHandler{
		cfg: cfg,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		authHeader := r.Header.Get("Authorization")
		tokenString := service.GetTokenFromBearerString(authHeader)

		claims, err := service.ValidateToken(tokenString, h.cfg.AccessTokenSecret)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		user, err := repo.NewUserRepository().GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		resp := response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		authHeader := r.Header.Get("Authorization")
		tokenString := service.GetTokenFromBearerString(authHeader)

		claims, err := service.ValidateToken(tokenString, h.cfg.RefreshTokenSecret)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := service.GenerateToken(claims.ID, h.cfg.AccessTokenLifetimeMinutes, h.cfg.AccessTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := service.GenerateToken(claims.ID, h.cfg.RefreshTokenLifetimeMinutes, h.cfg.RefreshTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := response.RefreshResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
