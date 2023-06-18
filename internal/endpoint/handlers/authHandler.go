package handler

import (
	"authJWT/configs"
	"authJWT/internal/db/model"
	"authJWT/internal/db/repo"
	"authJWT/internal/endpoint"
	response "authJWT/internal/endpoint/responses"
	"authJWT/internal/service"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	cfg *configs.Config
}

func NewAuthHandler(cfg *configs.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		req := new(endpoint.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := repo.NewUserRepository().GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := service.GenerateToken(user.ID, h.cfg.AccessTokenLifetimeMinutes, h.cfg.AccessTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := service.GenerateToken(user.ID, h.cfg.RefreshTokenLifetimeMinutes, h.cfg.RefreshTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := response.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		req := new(endpoint.SignUpRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Password != req.PasswordAgain {
			http.Error(w, "Password doesn't match!", http.StatusBadRequest)
			return
		}

		userRepo := repo.NewUserRepository()
		for _, user := range userRepo.Users {
			if user.Email == req.Email {
				http.Error(w, "Email already in use!", http.StatusBadRequest)
				return
			}
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		userRepo.Users = append(userRepo.Users, &model.User{
			Email:    req.Email,
			Password: string(password),
		})

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
