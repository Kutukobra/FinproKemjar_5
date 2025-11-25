package service

import (
	"context"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/model"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo     repository.UserRepository
	hashcost int
}

func NewUserService(repo repository.UserRepository, hashCost int) *UserService {
	return &UserService{repo: repo, hashcost: hashCost}
}

func (h *UserService) GetUser(ctx context.Context, username string) (*model.User, error) {
	user, err := h.repo.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *UserService) RegisterUser(ctx context.Context, username string, email string, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), h.hashcost)
	if err != nil {
		return nil, err
	}

	return h.repo.RegisterUser(ctx, username, email, string(hashedPassword))
}

func (h *UserService) LoginUser(ctx context.Context, username string, password string) (*model.User, error) {
	userData, err := h.repo.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	// Password comparison
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))

	if err != nil {
		return nil, err
	} else {
		return userData, nil
	}
}

func (h *UserService) ChangeUserPassword(ctx context.Context, username string, newPassword string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), h.hashcost)
	if err != nil {
		return nil, err
	}
	user, err := h.repo.ChangeUserPassword(ctx, username, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return user, nil
}
