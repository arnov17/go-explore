package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your local name")
	var user *string = flag.String("username", "Arnov", "put your username")
	var password *string = flag.String("password", "root", "Put your password")
	var number *int = flag.Int("number", 100, "put your number")

	flag.Parse()

	fmt.Println("host :", *host)
	fmt.Println("username :", *user)
	fmt.Println("password :", *password)
	fmt.Println("number :", *number)

}
