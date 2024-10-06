package users_services

import (
	"time"

	"github.com/montinger-com/montinger-server/app/shared/enums"
	users_model "github.com/montinger-com/montinger-server/app/users/models"
	users_repository "github.com/montinger-com/montinger-server/app/users/repositories"
	"github.com/montinger-com/montinger-server/lib/db"
	"github.com/montinger-com/montinger-server/lib/hashing"
)

type UsersService struct {
	userRepo *users_repository.UserRepository
}

func NewUserService() *UsersService {

	return &UsersService{
		userRepo: users_repository.NewUserRepository(db.MongoClient),
	}
}

func (s *UsersService) GetByEmail(email string) (*users_model.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *UsersService) Create(user *users_model.User) (*users_model.User, error) {
	hashedPassword, err := hashing.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userData := &users_model.User{
		Email:     user.Email,
		Password:  hashedPassword,
		Status:    enums.Active,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.userRepo.Create(userData); err != nil {
		return nil, err
	}

	return userData, nil
}
