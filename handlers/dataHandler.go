package dataHandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		}

		//Convert the JSON string into a struct
		var userLoginData LoginData
		if err := json.Unmarshal(data, &userLoginData); err != nil {
			fmt.Println(err)
		}

		//Validate whether a user has entered an existing email when trying to log in
		row, err := database.Query("SELECT * FROM users WHERE email = ?", userLoginData.Email)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(row)
		}
	}
}
