package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/montinger-com/montinger-server/app/utils"
	"github.com/rashintha/logger"
)

func GetJsonBody[BodyType any](c *gin.Context) BodyType {
	var data BodyType
	err := utils.AutoMapper(c.MustGet("body"), &data)
	if err != nil {
		logger.Errorln(err.Error())
	}
	return data
}
