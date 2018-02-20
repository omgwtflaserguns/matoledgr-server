package config

import (
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var logger = logging.MustGetLogger("log")

type ConfigLog struct {
	Level string
}

type ConfigGrpc struct {
	Address string
}

type ConfigGrpcWeb struct {
	Address string
}

type ConfigDatabase struct {
	File string
}

type Configuration struct {
	Log       ConfigLog
	Grpc      ConfigGrpc
	GrpcWeb   ConfigGrpcWeb
	Database  ConfigDatabase
	DebugCors bool
}

var config *Configuration

func GetConfig() Configuration {
	if config == nil {
		config = readConfig()
	}
	return *config
}

func readConfig() *Configuration {
	//TODO use config watch function and update values. See viper docs
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.matomat")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Panicf("Could not read Config: %s", err)
	}

	config := &Configuration{
		Log: ConfigLog{
			Level: viper.GetString("log.level"),
		},
		Grpc: ConfigGrpc{
			Address: viper.GetString("grpc.address"),
		},
		GrpcWeb: ConfigGrpcWeb{
			Address: viper.GetString("grpcWeb.address"),
		},
		Database: ConfigDatabase{
			File: viper.GetString("database.file"),
		},
		DebugCors: viper.GetBool("debugCors"),
	}

	logger.Debugf("config: %+v", config)
	return config
}
