package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("salah")

	if err == nil {
		fmt.Print(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(100000000, 10)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("200000")
	fmt.Println(valueInt)
}
