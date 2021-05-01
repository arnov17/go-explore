package main

import (
	"fmt"
)

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	// resultString := result.(string)
	// // resultString := result.(bool) // panic, karna beda type
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is String")
	case int:
		fmt.Println("value", value, "is integer")
	case bool:
		fmt.Println("value", value, "is Boolean")
	default:
		fmt.Println("unknown type")
	}

}
