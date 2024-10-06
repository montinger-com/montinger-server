package api

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"github.com/montinger-com/montinger-server/api/controllers"
	"github.com/montinger-com/montinger-server/app/utils/response"
	"github.com/montinger-com/montinger-server/app/utils/token"
	"github.com/montinger-com/montinger-server/config"
	cors_utils "github.com/montinger-com/montinger-server/lib/cors"
	"github.com/rashintha/logger"
)

var router *gin.Engine

func init() {
	if config.MODE == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()

	router.SetTrustedProxies(nil)
	router.Use(cors_utils.CORS())
	router.Use(requestid.New())
	router.Use(token.Interceptor())
	router.Use(response.Interceptor())
	router.Use(response.Log())

	controllers.Init(router)
}

func Run(addr ...string) {
	defer recoverOnPanic()
	router.Run(addr...)
}

func recoverOnPanic() {
	if r := recover(); r != nil {
		logger.Errorf("Recovered from panic: %v", r)
	}
}
