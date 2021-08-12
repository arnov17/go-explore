package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("input string")
	var inputWord string
	fmt.Scan(&inputWord)

	newWord := strings.ToLower(inputWord)

	if strings.HasPrefix(newWord, "i") && strings.HasSuffix(newWord, "n") && strings.Contains(newWord, "a") {
		fmt.Println("found!!")
	} else {
		fmt.Println("Not found!!")
	}

}
