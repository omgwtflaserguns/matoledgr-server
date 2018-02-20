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

func initializeDatabase() {
	logger.Debug("initializing db...")
	_, err := DbCon.Exec("CREATE TABLE Product (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"name VARCHAR(64), " +
		"price REAL" +
		");")

	if err != nil {
		logger.Panic(err)
	}

	logger.Debug("Inserting data...")
	_, err = DbCon.Exec("INSERT INTO Product (name, price)" +
		" VALUES ('Club Mate', 0.75);")

	if err != nil {
		logger.Panic(err)
	}
}

func openDatabase() {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Panic(err)
	}
	DbCon = db
}
