package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func userInput(prompt string) (string, error) {
	// User input
	fmt.Println(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return scanner.Text(), nil
}

func Swap(numSlice []int, index int) {
	temp := numSlice[index]
	numSlice[index] = numSlice[index+1]
	numSlice[index+1] = temp
}

func BubbleSort(numbers []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				swapped = true
			}
		}
	}

	// Print the sorted slice
	fmt.Println("Sorted:")
	fmt.Println(numbers)
}

func main() {
	numbers := make([]int, 0)

	for i := 0; i < 10; i++ {
		input, err := userInput("submit an integer:")
		if err != nil {
			panic(err)

		}

		number, err := strconv.ParseInt(input, 10, 0)
		if err != nil {
			panic(err)

		}

		numbers = append(numbers, int(number))
	}

	fmt.Println("submited integers:")
	fmt.Println(numbers)
	fmt.Println("------------")
	BubbleSort(numbers)
}
