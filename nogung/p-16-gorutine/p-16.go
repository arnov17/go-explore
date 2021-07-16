package main

import (
	"fmt"
	"time"
)

func main() {

	print(5, "halo")
	go print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
		time.Sleep(time.Second)
	}
}
