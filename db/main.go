package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var path = ""

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
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		fmt.Println("Database not found, creating...")
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
	fmt.Println("Creating Tables...")
	_, err := db.Exec("CREATE TABLE Product (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"name VARCHAR(64), " +
		"price REAL" +
		");")

	if err != nil {
		panic(err)
	}

	fmt.Println("Inserting Data...")
	_, err = db.Exec("INSERT INTO Product (name, price)" +
		" VALUES ('Club Mate', 0.75);")

	if err != nil {
		panic(err)
	}
}

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Failed DB open: %v", err)
		panic(err)
	}
	return db
}
