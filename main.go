package main

import (
	"github.com/montinger-com/montinger-server/config"
	"github.com/rashintha/logger"
)

func init() {
	logger.Defaultln("Starting server on " + config.HOST + ":" + config.PORT)
}

func main() {
	logger.Defaultln("Server started")
}
