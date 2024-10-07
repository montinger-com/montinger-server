package auth_services

import (
	"fmt"

	"github.com/montinger-com/montinger-server/app/shared/enums"
	users_repository "github.com/montinger-com/montinger-server/app/users/repositories"
	"github.com/montinger-com/montinger-server/app/utils/helpers"
	"github.com/montinger-com/montinger-server/lib/db"
	"github.com/montinger-com/montinger-server/lib/exceptions"
	"github.com/montinger-com/montinger-server/lib/hashing"
	jwt_utils "github.com/montinger-com/montinger-server/lib/jwt"
)

type AuthService struct {
	userRepo *users_repository.UsersRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: users_repository.NewUsersRepository(db.MongoClient),
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

	token, err := s.GetAuthorizedTokens(user.ID, user.Email, user.ID)
	return token, err
}

func (s *AuthService) GetAuthorizedTokens(id string, email string, alias string) (map[string]string, error) {

	var tokenPayload = make(map[string]string)

	tokenPayload[enums.AccessKey] = jwt_utils.GenerateAccessToken(jwt_utils.TokenPayload{
		ID:    id,
		Email: email,
		Alias: alias,
	})

	tokenPayload[enums.RefreshKey] = jwt_utils.GenerateRefreshToken(jwt_utils.TokenPayload{
		ID:    id,
		Email: email,
		Alias: alias,
	})

	return tokenPayload, nil
}

func (s *AuthService) GetRefreshTokenData(token string) (string, string, string, error) {

	verifiedToken, err := jwt_utils.ValidateRefreshToken(fmt.Sprintf("%v", token))
	if err != nil || !verifiedToken.Valid {
		return "", "", "", err
	}

	data, err := jwt_utils.GetTokenData(verifiedToken.Claims)
	if helpers.IsEmpty(data.Alias) || err != nil {
		return "", "", "", err
	}

	return data.ID, data.Email, data.Alias, nil
}
