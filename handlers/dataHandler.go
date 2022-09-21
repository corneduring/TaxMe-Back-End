package dataHandlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type UserData struct {
	Email             string
	Password          string
	PasswordValidator string
}

type Calculation struct {
	CalculationID    int
	UserID           int
	PaymentFrequency string `json:"Frequency"`
	Salary           string
	MonthlyTax       string
	YearlyTax        string
	Timestamp        time.Time
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func Login(database *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)

		//Convert the request into a readable JSON string
		requestString, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("Could not read your login data!"))
		}

		//Parse the JSON string into a struct
		var loginData UserData
		err = json.Unmarshal(requestString, &loginData)
		if err != nil {
			writer.Write([]byte("Could not read your login data!"))
		}

		//Prepare SQL statement
		stmt, err := database.Prepare("SELECT * FROM users WHERE email=$1")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		//Validate whether the email entered already exists
		result, err := stmt.Exec(loginData.Email)
		if err != nil {
			log.Print(err)
			writer.Write([]byte("An error has occured. Could not log you in."))
		}

		//Count the amount of rows retrieved from the database
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
		requestString, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("Couldn't read your data!"))
		}

		//Parse the JSON string into a struct
		var signUpData UserData
		err = json.Unmarshal(requestString, &signUpData)
		if err != nil {
			writer.Write([]byte("Couldn't read your data!"))
		}

		//Prepare SQL statement
		stmt, err := database.Prepare("SELECT * FROM users WHERE email = $1")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		//Validate whether the email entered already exists
		result, err := stmt.Exec(signUpData.Email)
		if err != nil {
			log.Print(err)
			writer.Write([]byte("Couldn't validate your email"))
		}

		//Count the amount of rows retrieved from the database
		valid, _ := result.RowsAffected()
		if valid == 1 {
			writer.Write([]byte("The email entered already exists!"))
			return
		}

		//Validate the syntax of the email entered
		validEmail, err := validateEmail(signUpData.Email)
		if !validEmail {
			writer.Write([]byte(err.Error()))
			return
		}

		//Validate the user's passwords to make it more secure
		validPassword, err := validatePassword(signUpData.Password)
		if !validPassword {
			writer.Write([]byte(err.Error()))
			return
		}

		//Validate whether the user's passwords match
		if signUpData.Password != signUpData.PasswordValidator {
			writer.Write([]byte("Your passwords don't match!"))
			return
		}

		_, err = database.Exec("INSERT INTO users(email, password) VALUES ($1, $2)", signUpData.Email, signUpData.Password)
	}
}

func GetHistory(database *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)

		//Convert the request into a readable JSON string
		requestString, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("Couldn't read your data!"))
		}

		//Parse the JSON string into a struct
		var loggedInUser UserData
		err = json.Unmarshal(requestString, &loggedInUser)
		if err != nil {
			writer.Write([]byte("Couldn't read your data!"))
		}

		stmt, err := database.Prepare("SELECT * FROM calculations WHERE user_id = (SELECT user_id FROM users WHERE email = $1);")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		calculations, err := stmt.Query(loggedInUser.Email)
		if err != nil {
			log.Print(err)
			writer.Write([]byte("An error has occured. Could not retrieve your calculations."))
		}

		var history []Calculation
		var response []byte
		defer calculations.Close()

		for calculations.Next() {
			var calculation Calculation

			if err := calculations.Scan(&calculation.CalculationID, &calculation.UserID, &calculation.Salary, &calculation.PaymentFrequency, &calculation.MonthlyTax, &calculation.YearlyTax, &calculation.Timestamp); err != nil {
				log.Print(err)
			}

			history = append(history, calculation)
		}
		response, err = json.MarshalIndent(history, "", "    ")
		if err != nil {
			log.Print(err)
			return
		}

		writer.Write(response)
	}
}

func validateEmail(email string) (bool, error) {
	valid, err := regexp.MatchString(`^([a-z|\d]+[\.|\-|_]?[a-z|\d]+)+@([a-z|\d]+\-?[a-z|\d]+)+\.[a-z]{2,63}$`, email)
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

func validatePassword(password string) (bool, error) {
	var valid bool
	var err error
	var upper, lower, number, special, length bool

	//Check if the password contains at least one upper case character
	if upper, _ = regexp.MatchString(`[A-Z]+`, password); !upper {
		err = errors.New("Your password must contain at least one upper case character")
	}

	//Check if the password contains at least one lower case character
	if lower, _ = regexp.MatchString(`[a-z]+`, password); !lower {
		err = errors.New("Your password must contain at least one lower case character")
	}

	//Check if the password contains at least one number
	if number, _ = regexp.MatchString(`\d+`, password); !number {
		err = errors.New("Your password must contain at least one number")
	}

	//Check if the password contains at least one special character
	if special, _ = regexp.MatchString(`[^\w\s]+`, password); !special {
		err = errors.New("Your password must contain at least one special character")
	}

	//Check if the password is at least 6 characters long
	if length = len(password) >= 6; !length {
		err = errors.New("Your password must be at least 6 characters long")
	}

	valid = upper && lower && number && special && length

	return valid, err
}
