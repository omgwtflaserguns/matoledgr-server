package db

import "github.com/omgwtflaserguns/matomat-server/util"

func initializeDatabase() {
	logger.Debug("initializing db...")

	initProduct()
	initAccount()
	initLogin()
	initPayment()
	initTransaction()
}

func initProduct() {
	logger.Debug("create table product")
	_, err := DbCon.Exec(
		"CREATE TABLE Product (" +
			"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
			"name VARCHAR(64) NOT NULL, " +
			"price REAL NOT NULL, " +
			"isActive INTEGER NOT NULL DEFAULT 1 " +
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
			"accountId INTEGER NOT NULL, " +
			"created TIMESTAMP NOT NULL, " +
			"FOREIGN KEY(accountId) REFERENCES Account(id) " +
			");")

	util.Check("Error creating table login: %v", err)
}

func initPayment() {
	logger.Debug("create table payment")
	_, err := DbCon.Exec(
		"CREATE TABLE Payment ( " +
			"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
			"accountId INTEGER NOT NULL, " +
			"value REAL NOT NULL, " +
			"timestamp TIMESTAMP NOT NULL, " +
			"FOREIGN KEY(accountId) REFERENCES Account(id) " +
			");")

	util.Check("Error creating table payment: %v", err)
}

func initTransaction() {
	logger.Debug("create table transaction")
	_, err := DbCon.Exec(
		"CREATE TABLE AccountTransaction ( " +
			"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
			"accountId INTEGER NOT NULL, " +
			"productId INTEGER NOT NULL, " +
			"paymentId INTEGER, " +
			"price REAL NOT NULL, " +
			"timestamp TIMESTAMP NOT NULL, " +
			"FOREIGN KEY(accountId) REFERENCES Account(id), " +
			"FOREIGN KEY(productId) REFERENCES Product(id) " +
			"FOREIGN KEY(paymentId) REFERENCES Payment(id) " +
			");")

	util.Check("Error creating table transaction: %v", err)
}
