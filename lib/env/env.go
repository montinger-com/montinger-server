package env

import (
	"os"
	"strings"

	"github.com/rashintha/logger"
)

var CONF = map[string]string{}

func init() {
	logger.Defaultln("Loading environment variables")
	data, err := os.ReadFile(".env")

	var sysVars []string
	if err != nil {
		logger.Warningf("%v", err.Error())
		sysVars = os.Environ()
	} else {
		sysVars = strings.FieldsFunc(string(data), split)
		sysVars = append(sysVars, os.Environ()...)
	}

	for _, val := range sysVars {
		if val[:1] != "#" {
			commentSplit := strings.Split(val, "#")
			commentLessString := strings.TrimSpace(commentSplit[0])

			split := strings.Split(commentLessString, "=")

			if len(split) < 2 {
				logger.ErrorFatal("Wrong format found in .env")
			}

			CONF[split[0]] = split[1]
		}
	}
	logger.Defaultf("%v", CONF)
}

func split(r rune) bool {
	return r == '\r' || r == '\n'
}
