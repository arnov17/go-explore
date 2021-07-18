package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
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

	// message := make(chan string)

	// for _, each := range []string{"wick", "hunt", "bourne"} {
	// 	go func(who string) {
	// 		var data = fmt.Sprintln("hello,", who)
	// 		message <- data
	// 		// time.Sleep(time.Second)
	// 	}(each)
	// }

	// for i := 0; i < 3; i++ {
	// 	printMessage(message)
	// }

	// // CHANEL BUFFER
	// pp := make(chan int, 3)

	// go func() {
	// 	for p := range pp {
	// 		p = <-pp
	// 		fmt.Println("Terima data ===========", p)
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("kirim data")
	// 	pp <- i
	// }

	// //CHANNE: SELECT // control penerimaan channel
	// runtime.GOMAXPROCS(2)
	// var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	// fmt.Println("numbers :", numbers)

	// var ch1 = make(chan float64)
	// go getAverage(numbers, ch1)

	// var ch2 = make(chan int)
	// go getMax(numbers, ch2)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case avg := <-ch1:
	// 		fmt.Printf("avg \t: %2f \n", avg)
	// 	case max := <-ch2:
	// 		fmt.Printf("max \t: %d \n", max)
	// 	}
	// }

	// CHANNEL CLOSE
	// runtime.GOMAXPROCS(2)

	// var messages = make(chan string)
	// go sendMessage(messages)
	// printMessage(messages)

	// CHANNEL TIMEOUT
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(5)

	var messages = make(chan int)

	go sendData(messages)
	retreiveData(messages)

}

// CHANNEL SELECT
func getAverage(number []int, ch chan float64) {
	var sum = 0
	for _, each := range number {
		sum += each
	}
	ch <- float64(sum) / float64(len(number))
}

func getMax(number []int, ch chan int) {
	var max = number[0]
	for _, e := range number {
		if max < e {
			max = e
		}
	}
	ch <- int(max)
}

// CHANNEL CLOSE
func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

// CHANNEL TIMEOUT
func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}

// func CetakNama(c chan string, v string) {
// 	c <- v
// }

// func printMessage(what chan string) {
// 	fmt.Println(<-what)
// }
