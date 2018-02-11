//go:generate protoc -I ./contracts --go_out=plugins=grpc:./generated ./contracts/matoledgr.proto

package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/omgwtflaserguns/matoledgr-server/db"
	"github.com/op/go-logging"
	"os"
	"sync"
)

var wg sync.WaitGroup
var logger = logging.MustGetLogger("log")
var leveledBackend logging.LeveledBackend

func main() {
	createLogger()
	readConfig()
	configureLogger()

	dbCon = db.Connect(config.database.file)
	grpcServer := createGrpcServer()
	wrapGrpcServer(grpcServer)

	logger.Debug("startup complete, listening...")
	wg.Wait()
}

func configureLogger() {
	var level logging.Level
	switch config.log.level {
	case "CRITICAL":
		level = logging.CRITICAL
	case "ERROR":
		level = logging.ERROR
	case "WARNING":
		level = logging.WARNING
	case "NOTICE":
		level = logging.NOTICE
	case "INFO":
		level = logging.INFO
	default:
		level = logging.DEBUG
	}
	logger.Debugf("Loglevel will now be set to %s", level)
	leveledBackend.SetLevel(logging.DEBUG, "")
}


func createLogger() {
	//TODO Implement file logger
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	backend := logging.NewLogBackend(os.Stdout, "", 0)
	formatedBackend := logging.NewBackendFormatter(backend, format)
	leveledBackend = logging.AddModuleLevel(formatedBackend)
	leveledBackend.SetLevel(logging.DEBUG, "")

	logging.SetBackend(leveledBackend)
}
