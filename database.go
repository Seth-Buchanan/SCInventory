package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"errors"
)
// Figure out how to open database once for program


func openDatabase(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)

	return db, err
}

type sqliteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *sqliteStore {
	return &sqliteStore{
		db: db,
	}
}


func (r *sqliteStore) Migrate() error {
	
	query := `
    CREATE TABLE IF NOT EXISTS ourtable (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        quantity INTEGER NOT NULL,
        replacement_level INTEGER
    );
    `
	_, err := r.db.Exec(query)
	return err
}


func (r *sqliteStore) add_quantity(id int, amount int) error {
	res, err := r.db.Exec("UPDATE ourtable SET quantity = quantity + ? where id = ?", amount, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Update Failed!")
	}

	return nil
}


func (r *sqliteStore) subtract_quantity(id int, amount int) error {
	res, err := r.db.Exec("UPDATE ourtable SET quantity = quantity - ? where id = ?", amount, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Update Failed!")
	}

	return nil
}

func (r *sqliteStore) item_quantity(id int) (int, string) {
	rows, err := r.db.Query("SELECT quantity, name FROM ourtable WHERE id = ?", id)
	if err != nil {
		log.Fatalf("ERROR could not query database %v", err)
	}
	defer rows.Close()
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


