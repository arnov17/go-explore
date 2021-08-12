package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	slice := make([]int, 0, 3)

	for {
		fmt.Println("Enter a number : ")
		fmt.Println("Submit 'X' to Exit")
		fmt.Scan(&input)

		if input == "X" || input == "x" {
			break
		} else {
			inputInt, _ := strconv.Atoi(input)

			slice = append(slice, inputInt)

			sort.Ints(slice)
			fmt.Println(slice)
		}

	}
}
