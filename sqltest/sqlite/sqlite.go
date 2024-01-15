package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

func main() {
	database, err := sql.Open("sqlite3", "./test.db")
	fmt.Printf("database variable is a %s\n", reflect.TypeOf(database))
	if err != nil {
		fmt.Printf("ERROR could not access database %v", err)
	}
	
	statement, err := database.Prepare( "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		fmt.Printf("ERROR could not create database %v", err)
	}
	
	statement.Exec()
	statement, err = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	if err != nil {
		fmt.Printf("ERROR could not insert into database %v", err)
	}
	
	statement.Exec("Seth", "Smith")
	rows, err := database.Query(fmt.Sprintf("SELECT id, firstname, lastname FROM people WHERE firstname = '%s'", "Seth"))
	if err != nil {
		fmt.Printf("ERROR could not query database %v", err)
	}
		
	var id int
	var firstname string
	var lastname string
	fmt.Printf("rows variable is a %s\n", reflect.TypeOf(rows))
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Printf("%d: firstname %s lastname %s\n", id, firstname, lastname )
	}
}
