package db

import (
	"database/sql"
	"os"
	"github.com/op/go-logging"
)

var path = ""
var logger = logging.MustGetLogger("log")

func Connect(databasePath string) *sql.DB {
	path = databasePath
	isNewDB := createIfNotFound()
	db := openDatabase()

	if isNewDB {
		initializeDatabase(db)
	}
	return db
}

func createIfNotFound() bool {
	logger.Debug("Searching databse in %s", path)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		logger.Debug("database not found, creating...")
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		file.Close()
		return true
	}
	return false
}

func initializeDatabase(db *sql.DB) {
	logger.Debug("initializing db...")
	_, err := db.Exec("CREATE TABLE Product (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"name VARCHAR(64), " +
		"price REAL" +
		");")

	if err != nil {
		panic(err)
	}

	logger.Debug("Inserting data...")
	_, err = db.Exec("INSERT INTO Product (name, price)" +
		" VALUES ('Club Mate', 0.75);")

	if err != nil {
		panic(err)
	}
}

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Fatalf("Failed to open db: %s %v", path, err)
		panic(err)
	}
	return db
}
