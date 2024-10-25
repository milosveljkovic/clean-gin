package service

import (
	"cleangin/internal/models"
	"cleangin/internal/repo"

	"golang.org/x/crypto/bcrypt"
)

type UserSvcI interface {
	RegisterUser(u models.User) error
}

type UserSvc struct {
	UserRepo repo.UserRepoI
}

func NewUserSvc(repo repo.UserRepoI) UserSvcI {
	return &UserSvc{UserRepo: repo}
}

func (svc *UserSvc) RegisterUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err = svc.UserRepo.RegisterUser(user); err != nil {
		return err
	}
	return nil
}
