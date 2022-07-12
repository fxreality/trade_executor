package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func returnDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/swissblock.db")
	checkErr(err)

	return db
}

func closeDB(db *sql.DB) {

	defer db.Close()
}

func createTables() {
	db := returnDB()
	stmt1, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS orders (id TEXT not null primary key, 
		asset_pair TEXT, order_type TEXT, price INT, size INT , created_at TEXT);`)
	stmt1.Exec()
	stmt2, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS trades (id TEXT not null primary key, order_id TEXT  not null, 
		asset_pair TEXT, order_type TEXT, price INT, size INT , executed_at TEXT, ob_update_id TEXT);`)
	stmt2.Exec()

	closeDB(db)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
