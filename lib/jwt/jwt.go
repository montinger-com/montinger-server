package jwt_utils

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/montinger-com/montinger-server/config"
)

type TokenPayload struct {
	ID          string `json:"id,omitempty" structs:"id"`
	Key         string `json:"key,omitempty" structs:"key"`
	Email       string `json:"email,omitempty" structs:"email"`
	Alias       string `json:"alias,omitempty" structs:"alias"`
	IsExchanged bool   `json:"is_exchanged,omitempty" structs:"is_exchanged"`
	ReadOnly    bool   `json:"read_only,omitempty" structs:"read_only"`
	Type        string `json:"type,omitempty" structs:"type"`
}

func GenerateAccessToken(payload TokenPayload) string {

	claims := &jwt.MapClaims{
		"iss":  config.ISSUER,
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.JWT_ACCESS_EXPIRES_IN_SECONDS))),
		"nbf":  jwt.NewNumericDate(time.Now()),
		"iat":  jwt.NewNumericDate(time.Now()),
		"data": GenerateDataMap(payload),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.JWT_ACCESS_SECRET))

	if err != nil {
		log.Fatalln("Generating JWT token failed.")
	}
	return signed
}

func GenerateRefreshToken(payload TokenPayload) string {

	claims := &jwt.MapClaims{
		"iss":  config.ISSUER,
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.JWT_REFRESH_EXPIRES_IN_SECONDS))),
		"nbf":  jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.JWT_ACCESS_EXPIRES_IN_SECONDS-5))),
		"iat":  jwt.NewNumericDate(time.Now()),
		"data": GenerateDataMap(payload),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.JWT_REFRESH_SECRET))

	if err != nil {
		log.Fatalln("Generating JWT refresh token failed.")
	}
	return signed
}

func GenerateDataMap(payload TokenPayload) map[string]interface{} {
	data := make(map[string]interface{})

	val := reflect.ValueOf(payload)
	typ := reflect.TypeOf(payload)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("structs")

		if !field.IsZero() {
			data[tag] = field.Interface()
		}
	}

	return data
}

func GetToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	return token
}

func ValidateAccessToken(tokenEncoded string) (*jwt.Token, error) {
	tk, err := jwt.Parse(tokenEncoded, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			if token.Header["alg"] != "HS256" {
				return nil, fmt.Errorf("invalid token signing algorithm: %v", token.Header["alg"])
			}
			return []byte(config.JWT_ACCESS_SECRET), nil
		}

		return nil, fmt.Errorf("invalid token signing method")
	})

	return tk, err
}
