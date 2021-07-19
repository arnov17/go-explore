package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// // #1
	// fmt.Println("start")
	// time.Sleep(time.Second * 4)
	// fmt.Println("after 4 seconds")

	// // #2
	// for true {
	// 	fmt.Println("Hello !!")
	// 	time.Sleep(1 * time.Second)
	// }

	// // #3
	// var timer = time.NewTimer(4 * time.Second)
	// fmt.Println("start")
	// <-timer.C
	// fmt.Println("finish")

	// // #4
	// var ch = make(chan bool)
	// time.AfterFunc(4*time.Second, func() {
	// 	fmt.Println("expired")
	// 	ch <- true
	// })

	// fmt.Println("start")
	// <-ch
	// fmt.Println("finsih")

	// // #5
	// <-time.After(4 * time.Second)
	// fmt.Println("expired")

	// // #6
	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)

	// func() {
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 		fmt.Println("Stop")
	// 		return
	// 	case t := <-ticker.C:
	// 		fmt.Println("hello", t)
	// 	}
	// }

	// #7
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Println("what is 725/25 ?")
	fmt.Scanln(&input)

	if input == "29" {
		fmt.Println("the answer is correct")
	} else {
		fmt.Println("the answer is wrong")
	}
}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}
