package main

import (
	"first_webapp/db"
	"first_webapp/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", dataHandlers.DataHandler)

	database := db.ConnectDatabase()
	defer database.Close()

	db.RunScript("db/db.sql", database)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Oops!")
	}
}
