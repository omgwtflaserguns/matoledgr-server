//go:generate protoc -I ./contracts --go_out=plugins=grpc:./generated ./contracts/matoledgr.proto

package main

import (
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"github.com/omgwtflaserguns/matoledgr-server/db"
	"github.com/op/go-logging"
	"os"
)

var wg sync.WaitGroup
var logger = logging.MustGetLogger("log")

func configureLogger() {
	//TODO Set Loglevel as command line argument
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	backend := logging.NewLogBackend(os.Stdout, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	module := logging.AddModuleLevel(backend)
	module.SetLevel(logging.DEBUG, "")

	logging.SetBackend(formatter)
}

func main() {
	configureLogger()
	logger.Debug("server starting")

	//TODO Read Database file as command line argument
	dbCon = db.Connect("./matoledgr.db", )

	logger.Debug("database connected")

	grpcServer, listener := createGrpcServer()

	wrapGrpcServer(grpcServer)

	logger.Debug("grpc-web started")

	wg.Add(1)
	go runGrpcServer(grpcServer, listener)

	logger.Debug("grpc started")
	logger.Debug("listening...")
	wg.Wait()
}
