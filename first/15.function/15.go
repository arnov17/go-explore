package main

import "fmt"

func sayHello() {
	fmt.Println("say hello")
}

func fullname(firstName string, lastName string) {
	fmt.Println(firstName, lastName)

}

// return value
func getHello(name string) string {
	if name == "" {
		return "nama kosong"
	} else {
		return "Hai " + name
	}
}

// return multiple value
func getFullName(name string) (string, int, string, int) {
	return "The", 1, name, 1994
}

func main() {

	for i := 1; i <= 10; i++ {
		sayHello()
	}
	firstName := "Arnov"
	fullname(firstName, "julian")

	// func return value
	result := getHello("Ani")
	fmt.Println(result)
	fmt.Println(getHello(""))

	//func return multiple value
	// _ menghiraukan value
	first, number, name, _ := getFullName("kemosabe")
	fmt.Println(first + " first")
	fmt.Println(number)
	fmt.Println(name + " name")

}
