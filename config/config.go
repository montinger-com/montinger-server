package config

import (
	"strconv"

	"github.com/montinger-com/montinger-server/lib/env"
)

var HOST string
var PORT string

var APP_ENV string
var MODE string

var DB_TYPE string
var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASS string
var DB_NAME string

var ISSUER string
var JWT_ACCESS_EXPIRES_IN_SECONDS int
var JWT_ACCESS_SECRET string
var JWT_REFRESH_EXPIRES_IN_SECONDS int
var JWT_REFRESH_SECRET string

func init() {
	HOST = env.CONF["HOST"]
	PORT = env.CONF["PORT"]

	APP_ENV = env.CONF["APP_ENV"]
	MODE = env.CONF["MODE"]

	DB_TYPE = env.CONF["DB_TYPE"]
	DB_HOST = env.CONF["DB_HOST"]
	DB_PORT = env.CONF["DB_PORT"]
	DB_USER = env.CONF["DB_USER"]
	DB_PASS = env.CONF["DB_PASS"]
	DB_NAME = env.CONF["DB_NAME"]

	ISSUER = env.CONF["ISSUER"]
	JWT_ACCESS_SECRET = env.CONF["JWT_ACCESS_SECRET"]
	JWT_REFRESH_SECRET = env.CONF["JWT_REFRESH_SECRET"]

	expiresIn, errExp := strconv.Atoi(env.CONF["JWT_ACCESS_EXPIRES_IN_SECONDS"])
	if errExp != nil {
		expiresIn = 60
	}
	JWT_ACCESS_EXPIRES_IN_SECONDS = expiresIn

	expiresIn, errExp = strconv.Atoi(env.CONF["JWT_REFRESH_EXPIRES_IN_SECONDS"])
	if errExp != nil {
		expiresIn = 60
	}
	JWT_REFRESH_EXPIRES_IN_SECONDS = expiresIn
}
