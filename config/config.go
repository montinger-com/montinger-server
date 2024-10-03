package config

import "github.com/montinger-com/montinger-server/lib/env"

var HOST string
var PORT string

func init() {
	HOST = env.CONF["HOST"]
	PORT = env.CONF["PORT"]
}
