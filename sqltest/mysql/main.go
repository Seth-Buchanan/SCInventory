package main

import (
	"database/sql" 
	"fmt"
	_"github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	fmt.Println("Drivers: ", sql.Drivers())
	db, err := sql.Open("mysql", "root:<password-here>@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal("Could not connect to db");
	}
	defer db.Close()
	results, err := db.Query("select * from product");

	if err != nil {
		log.Fatal("Error when fetching product table rows: ", err);
	}

	defer results.Close()// Iterating through rows of a column
	for results.Next() {
		var (
			id int
			name string
			price int
		)
		err = results.Scan(&id, &name, &price) 
		if err != nil {
			log.Fatal("Unable to parse row:", err)
		}
		fmt.Printf("ID: %d, Name: '%s', Price: %d\n", id, name, price);
	}
	
	var (
		id int
		name string
		price int
	)
	err = db.QueryRow("Select * from product where id = 1").Scan(&id, &name, &price) // Specific item
	if err != nil {
		log.Fatal("Unable to parse row: ", err)
	}
	fmt.Printf("ID: %d, Name: '%s', Price: %d\n", id, name, price);

	// Adding prepared statements
	products := []struct {
		name string
		price int
	}{
		{"Light", 10},
		{"Mic", 30},
		{"Router", 90},
	}
	stmt, err := db.Prepare("INSERT INTO product (name, price) VALUES (?, ?)")
	defer stmt.Close()
	
	if err != nil {
		log.Fatal("Unable to prepare statement: ", err);
	}
	for _, product := range products {
		_, err = stmt.Exec(product.name, product.price)
		if err != nil {
			log.Fatal("Unable to execute statement: ", err)
		}
	}
	
	
}
