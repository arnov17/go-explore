package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("hallo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) chaneName2(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func main() {
	var s1 = student{"John Wick", 27}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName", s1.name)

	s1.chaneName2("ethant hunt")
	fmt.Println("s1 after changeName2", s1.name)

	var s2 = &student{"ethan hunt", 22}
	s2.sayHello()

}
