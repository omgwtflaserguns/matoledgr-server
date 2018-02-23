package db

import (
	"database/sql"
	"github.com/op/go-logging"
	"os"
)

var path = ""
var logger = logging.MustGetLogger("log")
var DbCon *sql.DB

func Connect(databasePath string) {
	logger.Debugf("connecting to database %s", databasePath)
	path = databasePath
	isNewDB := createIfNotFound()
	openDatabase()
	if isNewDB {
		initializeDatabase()
	}
	logger.Debug("database connected")
}

func Close() {
	DbCon.Close()
	logger.Debug("database closed")
}

func createIfNotFound() bool {
	logger.Debug("Searching databse in %s", path)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		logger.Debug("database not found, creating...")
		file, err := os.Create(path)
		if err != nil {
			logger.Panic(err)
		}
		file.Close()
		return true
	}
	return false
}

func openDatabase() {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Panic(err)
	}
	DbCon = db
}
