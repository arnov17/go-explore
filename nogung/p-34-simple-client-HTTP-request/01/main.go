package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var baseURL = "http://localhost:8080"

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	return data, nil

}

func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}

}
