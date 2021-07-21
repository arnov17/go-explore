package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1)
	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])
}
