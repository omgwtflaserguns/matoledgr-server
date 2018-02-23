package db

import "github.com/omgwtflaserguns/matomat-server/util"

func initializeDatabase() {
	logger.Debug("initializing db...")

	initProduct()
	initAccount()
	initLogin()
}

func initProduct() {
	logger.Debug("create table product")
	_, err := DbCon.Exec(
		"CREATE TABLE Product (" +
			"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
			"name VARCHAR(64) NOT NULL, " +
			"price REAL NOT NULL" +
			");")

	util.Check("Error creating table product: %v", err)

	logger.Debug("Inserting products")
	_, err = DbCon.Exec("INSERT INTO Product (name, price)" +
		" VALUES ('Club Mate', 0.75);")

	util.Check("Error inserting products: %v", err)
}

func initAccount() {
	logger.Debug("create table account")
	_, err := DbCon.Exec(
		"CREATE TABLE Account (" +
			"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
			"username VARCHAR(64) NOT NULL, " +
			"hash VARCHAR(60) COLLATE BINARY NOT NULL" +
			");")

	util.Check("Error creating table account: %v", err)
}

func initLogin() {
	logger.Debug("create table login")
	_, err := DbCon.Exec(
		"CREATE TABLE Login (" +
			"cookie VARCHAR(128) PRIMARY KEY, " +
			"accountId INTEGER, " +
			"created TIMESTAMP NOT NULL, " +
			"FOREIGN KEY(accountId) REFERENCES Account(id) " +
			");")

	util.Check("Error creating table login: %v", err)
}
