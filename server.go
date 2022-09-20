package main

import (
	"first_webapp/db"
	"first_webapp/handlers"
	"log"
	"net/http"
)

func main() {
	database := db.ConnectDatabase()

	http.HandleFunc("/login", dataHandlers.Login(database))
	http.HandleFunc("/signup", dataHandlers.SignUp(database))
	http.HandleFunc("/history", dataHandlers.GetHistory(database))

	defer database.Close()

	db.RunScript("db/db.sql", database)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
