package auth_services

import (
	"github.com/montinger-com/montinger-server/app/shared/enums"
	users_repository "github.com/montinger-com/montinger-server/app/users/repositories"
	"github.com/montinger-com/montinger-server/lib/db"
	"github.com/montinger-com/montinger-server/lib/exceptions"
	"github.com/montinger-com/montinger-server/lib/hashing"
	jwt_utils "github.com/montinger-com/montinger-server/lib/jwt"
)

type AuthService struct {
	userRepo *users_repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: users_repository.NewUserRepository(db.MongoClient),
	}
}

func (s *AuthService) AuthenticateUser(email string, password string) (map[string]string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, exceptions.UserNotFoundInOurDB
	}

	isValid, _ := hashing.VerifyPassword(password, user.Password)

	if !isValid {
		return nil, exceptions.InvalidUsernameOrPassword
	}

	token, err := s.GetAuthorizedTokens(user.ID, user.Email)
	return token, err
}

func (s *AuthService) GetAuthorizedTokens(id string, email string) (map[string]string, error) {

	var tokenPayload = make(map[string]string)

	tokenPayload[enums.AccessKey] = jwt_utils.GenerateAccessToken(jwt_utils.TokenPayload{
		ID:    id,
		Email: email,
		Alias: id,
	})

	tokenPayload[enums.RefreshKey] = jwt_utils.GenerateRefreshToken(jwt_utils.TokenPayload{
		ID:    id,
		Email: email,
		Alias: id,
	})

	return tokenPayload, nil
}
