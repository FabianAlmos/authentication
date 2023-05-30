package repo

import (
	"authJWT/internal/db/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Users []*model.User
}

func NewUserRepository() *UserRepository {
	p1, _ := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("22222222"), bcrypt.DefaultCost)

	users := []*model.User{
		{ID: 1, Name: "Alex", Email: "alex@example.com", Password: string(p1)},
		{ID: 2, Name: "Mary", Email: "mary@example.com", Password: string(p2)},
	}
	return &UserRepository{
		Users: users,
	}
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	for _, u := range r.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	for _, u := range r.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
