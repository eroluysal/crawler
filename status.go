package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Status struct {
	Status    string `json:"status"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type Response struct {
	Statuses []Status
}

func main() {
	response, err := http.Get("https://status.github.com/api/messages.json")
	exitWithErrorMessage(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	exitWithErrorMessage(err)

	values := Response{}.Statuses
	json.Unmarshal([]byte(body), &values)

	fmt.Println(values)
}

func exitWithErrorMessage(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
