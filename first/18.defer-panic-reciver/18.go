package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("error dengan messgae", message)
	}
	fmt.Println("applikasi selesai")
}

func callName() string {
	fmt.Println("testing callname")
	return "hello world"
}

func runApp(error bool) {
	defer endApp()
	defer callName()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("applikasi berjalan")

}

func main() {
	runApp(true)
	fmt.Println("main berjalan")

}

// output nya
// error dengan messgae APLIKASI ERROR
// applikasi selesai
// testing callname
// main berjalan
