package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname", name)
	} else {
		fmt.Println(err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	//lebih lengkap pakcage os bisa dilihat di golang.org/pkg/os
}
