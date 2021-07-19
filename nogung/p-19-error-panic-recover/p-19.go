package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// defer catch()

	// var input string

	// fmt.Println("type some number")
	// fmt.Scanln(&input)

	// var number int
	// var err error
	// number, err = strconv.Atoi(input)

	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(input, "is Number")
	// 	fmt.Println(err.Error())
	// }

	// //Custom Error
	// var name string
	// fmt.Println("type your name")
	// fmt.Scanln(&name)

	// if valid, err := validate(name); valid {
	// 	fmt.Println("halo", name)
	// } else {
	// 	// fmt.Println(err.Error())
	// 	panic(err.Error()) // mengunakan panic
	// }

	// recover IIFE
	data := []string{"superman", "aquaman", "wonder woman"}
	for _, each := range data {

		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("application Runn Perfecly")
				}
			}()
			panic("some error happen")
		}()
	}
}

// Custom Error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot Be Empty")
	}
	return true, nil
}

// recover
func catch() {
	// u/ handle panic
	if r := recover(); r != nil {
		fmt.Println("Error Occured", r)
	} else {
		fmt.Println("App running perfecly")
	}
}
