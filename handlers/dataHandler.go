package dataHandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginData struct {
	Email    string
	Password string
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func Login(database *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)

		//Convert the JSON data from the request into a readable string
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err)
			writer.Write([]byte("Could not read your login data!"))
		}

		//Convert the JSON string into a struct
		var userLoginData LoginData
		if err := json.Unmarshal(data, &userLoginData); err != nil {
			fmt.Println(err)
			writer.Write([]byte("Could not read your login data!"))
		}

		//Validate whether a user has entered an existing email when trying to log in
		result, err := database.Exec("SELECT * FROM users WHERE email=$1", userLoginData.Email)
		if err != nil {
			fmt.Println(err)
		}

		valid, _ := result.RowsAffected()
		if valid != 1 {
			log.Printf("The email '%s' does not exist!", userLoginData.Email)
			writer.Write([]byte("The email you entered does not exist. Try signing up for an account."))
			return
		}

		//Validate whether the password entered by the user matches the corresponding email in the database
		result, err = database.Exec("SELECT * FROM users WHERE email=$1 AND password=$2", userLoginData.Email, userLoginData.Password)
		if err != nil {
			fmt.Println(err)
		}

		valid, _ = result.RowsAffected()
		if valid != 1 {
			log.Println("The password does not match the email entered!", userLoginData.Email)
			writer.Write([]byte("Your password is invalid!"))
			return
		}
	}
}
