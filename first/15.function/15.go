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

// return named value
func storeFullname() (firstName, middlename, lastname string) {
	firstName = "ani"
	middlename = "marathus"
	lastname = "sholikah"
	return
}

//varidic func
// bisa lebih dari 1 param, tapi diakhir
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
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

	//functon named value
	fmt.Println(storeFullname())
	a, b, c := storeFullname()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// variadic func
	total := sumAll(10, 10, 30, 30)
	fmt.Println(total)

	slice := []int{10, 30, 40}
	totalslice := sumAll(slice...)
	fmt.Println(totalslice)

	// func as variable
	varGetHello := getHello("anna")
	fmt.Println(varGetHello)

}
