package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func sayHi(user Customer, name string) {
	fmt.Println("Hello", name, "My Name is", user.Name)
}

//struct method
func (user Customer) sayHi2(name string) {
	fmt.Println("Hello", name, "My Name is", user.Name)
}

func main() {
	var arnov Customer
	arnov.Name = "Arnov"
	arnov.Address = "Jakarta"
	arnov.Age = 26

	fmt.Println(arnov)

	ani := Customer{
		Name:    "ani",
		Address: "Malang",
		Age:     27,
	}

	fmt.Println(ani)

	anna := Customer{"Anna", "Depok", 22}
	fmt.Println(anna)

	sayHi(arnov, "zoro")

	//struct method
	arnov.sayHi2("Lufy")

}
