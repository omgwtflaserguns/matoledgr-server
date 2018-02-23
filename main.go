//go:generate protoc -I ./contracts --go_out=plugins=grpc:./generated ./contracts/matomat.proto

package main

import (
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/omgwtflaserguns/matomat-server/config"
	"github.com/omgwtflaserguns/matomat-server/db"
	"github.com/omgwtflaserguns/matomat-server/service"
	"github.com/op/go-logging"
	"math/rand"
	"time"
)

var wg *sync.WaitGroup
var logger = logging.MustGetLogger("log")
var leveledBackend logging.LeveledBackend
var conf config.Configuration

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	createLogger()
	conf = config.GetConfig()
	configureLogger()

	db.Connect(conf.Database.File)
	defer db.Close()

	wg = &sync.WaitGroup{}
	grpcServer := service.CreateGrpcServer(wg)
	service.WrapGrpcServer(grpcServer, wg)

	logger.Debug("startup complete, listening...")
	wg.Wait()
}

func configureLogger() {
	var level logging.Level
	switch conf.Log.Level {
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
