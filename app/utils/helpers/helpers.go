package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/montinger-com/montinger-server/app/shared/enums"
	"github.com/montinger-com/montinger-server/lib/utilities"
	"github.com/rashintha/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetJsonBody[BodyType any](c *gin.Context) BodyType {
	var data BodyType
	err := utilities.AutoMapper(c.MustGet("body"), &data)
	if err != nil {
		logger.Errorln(err.Error())
	}
	return data
}

func IsEmpty(str string) bool {
	return len(str) == 0 && str != enums.Null
}

func ObjectIDToString(s interface{}) string {
	stringId := s.(primitive.ObjectID).Hex()
	return stringId
}

func GetParams[BodyType any](c *gin.Context) BodyType {
	var data BodyType
	err := utilities.AutoMapper(c.MustGet("params"), &data)
	if err != nil {
		logger.Errorln(err.Error())
	}
	return data
}

func GetAuthToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	return token
}

func GetAPIKey(c *gin.Context) string {
	authHeader := c.GetHeader("X-API-Key")
	return authHeader
}
