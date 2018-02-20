package util

import "github.com/op/go-logging"

var logger = logging.MustGetLogger("log")

func Check(msg string, err error) {
	if err != nil {
		logger.Panic(msg, err)
	}
}
