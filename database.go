package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)
// Figure out how to open database once for program

func ammo_quantity(identifier int) (int, string) {
	databaseName := "./database.db"
	database, err := sql.Open("sqlite3", databaseName)
	
	if err != nil {
		log.Fatalf("ERROR could not access %s database: %v", databaseName, err)
	}

	rows, err := database.Query("SELECT quantity, name FROM ammunition WHERE identifier = ?", identifier)
	defer rows.Close()
	if err != nil {
		log.Fatalf("ERROR could not query database %v", err)
	}
	for rows.Next() {
		var quantity int
		var name string
		if err := rows.Scan(&quantity, &name); err != nil {
			log.Fatal(err)
		} else {
			return quantity, name
		}
	}

	return -1, ""
}


func add_quantity(identifier int, amount int) {
	databaseName := "./database.db"
	database, err := sql.Open("sqlite3", databaseName)
	
	if err != nil {
		log.Fatalf("ERROR could not access %s database: %v", databaseName, err)
	}

	statement, err := database.Prepare("UPDATE ammunition SET quantity = quantity + ? where identifier = ?", )
	defer statement.Close()
	if err != nil {
		log.Fatal("Unable to prepare statement ", err)
	}
	statement.Exec(amount, identifier)
}


func subtract_quantity(identifier int, amount int) {
	databaseName := "./database.db"
	database, err := sql.Open("sqlite3", databaseName)
	
	if err != nil {
		log.Fatalf("ERROR could not access %s database: %v", databaseName, err)
	}

	statement, err := database.Prepare("UPDATE ammunition SET quantity = quantity - ? where identifier = ?", )
	defer statement.Close()
	if err != nil {
		log.Fatal("Unable to prepare statement ", err)
	}
	statement.Exec(amount, identifier)
}
