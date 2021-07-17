package main

import (
	"fmt"
)

func main() {

	// c := make(chan string)

	// go func(n string) {
	// 	c <- n
	// }("Arnov")

	// var chanFunc = func(n string) {
	// 	c <- n
	// }

	// go chanFunc("testing")

	// b := <-c
	// d := <-c
	// fmt.Println(b)
	// fmt.Println(d)

	// var messages = make(chan string)

	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }

	// go sayHelloTo("john wick")
	// go sayHelloTo("ethan hunt")
	// go sayHelloTo("jason bourne")

	// var message1 = <-messages
	// fmt.Println(message1)

	// var message2 = <-messages
	// fmt.Println(message2)

	// var message3 = <-messages
	// fmt.Println(message3)

	// checkDua := make(chan string)
	// go CetakNama(checkDua, "yohoho")

	// result := <-checkDua
	// fmt.Println(result)

	// runtime.GOMAXPROCS(2)
	message := make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintln("hello,", who)
			message <- data
			// time.Sleep(time.Second)
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(message)
	}
}

func CetakNama(c chan string, v string) {
	c <- v
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
