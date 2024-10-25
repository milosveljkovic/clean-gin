package repo

import (
	"cleangin/internal/models"

	"gorm.io/gorm"
)

type UserRepoI interface {
	RegisterUser(u models.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoI {
	return &UserRepo{db: db}
}

func (repo *UserRepo) RegisterUser(u models.User) error {
	result := repo.db.Exec(`
	INSERT INTO public.users(name, email, password)
	VALUES (@Name, @Email, @Password);`, u)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
