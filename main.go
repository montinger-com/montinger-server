package main

import (
	"github.com/montinger-com/montinger-server/api"
	"github.com/montinger-com/montinger-server/config"
	"github.com/rashintha/logger"
)

func init() {
	logger.Defaultln("Starting server on " + config.HOST + ":" + config.PORT)
}

func main() {
	api.Run(config.HOST + ":" + config.PORT)
}
