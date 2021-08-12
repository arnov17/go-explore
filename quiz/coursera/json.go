package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	var address string
	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Enter address:")
	fmt.Scan(&address)

	person := map[string]string{
		"name":    name,
		"address": address,
	}

	result, _ := json.Marshal(person)
	fmt.Println(string(result))
}
