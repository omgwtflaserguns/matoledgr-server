package main

import (
	"github.com/spf13/viper"
)

var config Config

type ConfigLog struct {
	level     string
	writeFile bool
	file      string
}

type ConfigGrpc struct {
	address string
}

type ConfigGrpcWeb struct {
	address string
}

type ConfigDatabase struct {
	file string
}

type Config struct {
	log       ConfigLog
	grpc      ConfigGrpc
	grpcWeb   ConfigGrpcWeb
	database  ConfigDatabase
	debugCors bool
}

func readConfig() {
	//TODO use config watch function and update values. See viper docs
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.matoledgr")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Panicf("Could not read Config: %s", err)
	}

	config = Config{
		log: ConfigLog{
			level:     viper.GetString("log.level"),
			writeFile: viper.GetBool("log.writeFile"),
			file:      viper.GetString("log.file"),
		},
		grpc: ConfigGrpc{
			address: viper.GetString("grpc.address"),
		},
		grpcWeb: ConfigGrpcWeb{
			address: viper.GetString("grpcWeb.address"),
		},
		database: ConfigDatabase{
			file: viper.GetString("database.file"),
		},
		debugCors: viper.GetBool("debugCors"),
	}

	logger.Debugf("config: %+v", config)
}
