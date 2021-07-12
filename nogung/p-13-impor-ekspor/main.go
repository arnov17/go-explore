package main

import (
	f "fmt"
	"go-explore/nogung/p-13-impor-ekspor/helper"
	stu "go-explore/nogung/p-13-impor-ekspor/library"
)

func main() {
	var s1 = stu.Student{"ethan", 66}
	f.Println("name ", s1.Name)
	f.Println("grade", s1.Grade)
	// library.SayHello()
	helper.CallNickName()
}
