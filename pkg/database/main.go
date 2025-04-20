package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDataBase() error {
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS passwords (service TEXT PRIMARY KEY, password TEXT)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func AddPassWord(service string, password string) error {
	if db == nil {
		return errors.New("Database connection is nil")
	}

	statement, err := db.Prepare("INSERT INTO passwords(service, password) VALUES(?, ?)")
	if err != nil {
		log.Println("Error in preparing statement:", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(service, password)
	if err != nil {
		log.Println("Error in executing statement:", err)
		return err
	}

	return nil
}

func GetPassWord(service string) (string, error) {
	if db == nil {
		return "", errors.New("Database connection is nil")
	}

	statement, err := db.Prepare("SELECT password FROM passwords WHERE service=?")
	if err != nil {
		log.Println("Error in preparing statement:", err)
		return "", err
	}
	defer statement.Close()

	row := statement.QueryRow(service)
	var password string
	err = row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("String is not found")
		}
		log.Println("Error in string scan:", err)
		return "", err
	}

	return password, nil
}

func ClearPassWord(service string) error {
	if db == nil {
		return errors.New("Database connection is nil")
	}
	statement, err := db.Prepare("DELETE FROM passwords WHERE service=?")
	if err != nil {
		log.Println(err)
	}
	defer statement.Close()
	_, err = statement.Exec(service)
	if err != nil {
		log.Println(err)
	}
	return nil
}
