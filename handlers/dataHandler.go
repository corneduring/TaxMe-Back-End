package dataHandlers

import (
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

func DataHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", data)

	var userLoginData LoginData
	if err := json.Unmarshal(data, &userLoginData); err != nil {
		fmt.Println(err)
	}
}
