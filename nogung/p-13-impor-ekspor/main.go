package main

import (
	f "fmt"
	"go-explore/nogung/p-13-impor-ekspor/helper"
	"go-explore/nogung/p-13-impor-ekspor/library"
)

// "go-explore/nogung/p-13-impor-ekspor/helper"
// stu "go-explore/nogung/p-13-impor-ekspor/library"

func main() {
	// var s1 = library.Student{"ethan", 66}
	// f.Println("name ", s1.Name)
	// f.Println("grade", s1.Grade)
	// library.SayHello()
	helper.CallNickName()

	sayHello("ethan")

	f.Printf("Name  : %s\n", library.Student.Name)
	f.Printf("Grade : %d\n", library.Student.Grade)

}
