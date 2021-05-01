package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// interface kosong
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var eko Person
	eko.Name = "eko"

	SayHello(eko)

	cat := Animal{
		Name: "rokes",
	}

	SayHello(cat)

	// interface kosong
	var data interface{} = Ups(1)
	fmt.Println(data)
}
