package utils

import (
	"database/sql"
	"log"
)

//This is going to be the file that main calls to launch up the SQL connection

var user = "root"
var password = "password"
var host = "localhost"
var port = "3306"
var database = "golang"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		return nil, err
	}

	// Use a deferred function to ensure that the database connection is closed
	// before the function returns.
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		}
	}()

	// Perform a ping to verify that the database connection is valid.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
