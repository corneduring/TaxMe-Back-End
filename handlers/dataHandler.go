package dataHandlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type LoginData struct {
	Email    string
	Password string
}

type SignUpData struct {
	Email     string
	Password1 string
	Password2 string
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func Login(database *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)

		//Convert the request into a readable JSON string
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("Could not read your login data!"))
		}

		//Parse the JSON string into a struct
		var loginData LoginData
		if err := json.Unmarshal(data, &loginData); err != nil {
			writer.Write([]byte("Could not read your login data!"))
		}

		//Validate whether the email entered already exists
		result, err := database.Exec("SELECT * FROM users WHERE email=$1", loginData.Email)
		if err != nil {
			log.Print(err)
		}

		valid, _ := result.RowsAffected()
		if valid != 1 {
			writer.Write([]byte("The email you entered does not exist. Try signing up for an account first."))
			return
		}

		//Validate whether the password entered by the user matches the corresponding email in the database
		result, err = database.Exec("SELECT * FROM users WHERE email=$1 AND password=$2", loginData.Email, loginData.Password)
		if err != nil {
			log.Print(err)
		}

		valid, _ = result.RowsAffected()
		if valid != 1 {
			writer.Write([]byte("Your password is invalid!"))
			return
		}
	}
}

func SignUp(database *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)

		//Convert the request into a readable JSON string
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("Couldn't read your data!"))
		}

		//Parse the JSON string into a struct
		var signUpData SignUpData
		err = json.Unmarshal(data, &signUpData)
		if err != nil {
			writer.Write([]byte("Couldn't read your data!"))
		}

		//Validate whether the email entered already exists
		result, err := database.Exec("SELECT * FROM users WHERE email = $1", signUpData.Email)
		if err != nil {
			log.Print(err)
		}

		valid, _ := result.RowsAffected()
		if valid == 1 {
			writer.Write([]byte("The email entered already exists!"))
			return
		}

		//Validate the syntax of the email entered
		validEmail, err := isEmail(signUpData.Email)
		if !validEmail {
			writer.Write([]byte(err.Error()))
			return
		}

		//Validate the user's passwords to make it more secure
		validPassword, err := isPassword(signUpData.Password1)
		if !validPassword {
			writer.Write([]byte(err.Error()))
			return
		}

		//Validate whether the user's passwords match
		if signUpData.Password1 != signUpData.Password2 {
			writer.Write([]byte("Your passwords don't match!"))
			return
		}

		_, err = database.Exec("INSERT INTO users VALUES ($1, $2)", signUpData.Email, signUpData.Password1)
	}
}

func isEmail(email string) (bool, error) {
	valid, err := regexp.MatchString("^([a-z|\\d]+[\\.|\\-|_]?[a-z|\\d]+)+@([a-z|\\d]+\\-?[a-z|\\d]+)+\\.[a-z]{2}$", email)
	if err != nil {
		log.Println(err)
		return false, errors.New("Could not validate email!")
	}

	if !valid {
		return false, errors.New("You did not enter a valid email address!")
	} else {
		if len(email) > 254 {
			return false, errors.New("Your email is too long! Maximum email length is 254 characters.")
		}
	}

	return true, nil
}

func isPassword(password string) (bool, error) {
	valid, err := regexp.MatchString("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,10}$", password)
	if err != nil {
		log.Println(err)
		return false, errors.New("Could not validate password!")
	}

	if !valid {
		return false, errors.New(
			"Your password must match the following criteria:" +
				"Minimum of 8 characters" +
				"Maximum of 40 characters" +
				"At least one uppercase letter" +
				"At least one lower case letter" +
				"At least one number" +
				"At least one special character")
	}

	return true, nil
}
