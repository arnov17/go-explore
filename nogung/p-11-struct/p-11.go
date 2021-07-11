package main

import "fmt"

func main() {
	// var s1 student
	// s1.name = "John wick"
	// s1.grade = 2

	// fmt.Println("name  :", s1.name)
	// fmt.Println("grade :", s1.grade)
	// fmt.Println(s1)

	// var s1 = student{}
	// s1.name = "wick"
	// s1.grade = 2

	// var s2 = student{"ethan", 2}

	// var s3 = student{name: "jason"}

	// var s4 = student{name: "wayne", grade: 2}
	// var s5 = student{grade: 2, name: "bruce"}

	// fmt.Println("student 1 :", s1.name)
	// fmt.Println("student 2 :", s2.name)
	// fmt.Println("student 3 :", s3.name)
	// fmt.Println("student 1 :", s1)
	// fmt.Println("student 2 :", s2)
	// fmt.Println("student 3 :", s3)
	// fmt.Println("student 3 :", s4)
	// fmt.Println("student 3 :", s5)

	// // Variabel Objek Pointer
	// var s1 = student{name: "wick", grade: 2}

	// var s2 *student = &s1
	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 4, name :", s2.name)

	// s2.name = "ethan"
	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 4, name :", s2.name)

	// // embed struct
	// var s1 = student{}
	// s1.name = "wick"
	// s1.age = 21        // age of student
	// s1.person.age = 22 // age of person

	// fmt.Println(s1.name)
	// fmt.Println(s1.age)
	// fmt.Println(s1.person.age)

	// // sub-struct
	// var p1 = person{name: "wick", age: 21}
	// var s1 = student{person: p1, grade: 2}

	// fmt.Println("name  :", s1.name)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("age   :", s1.age)
	// fmt.Println("grade :", s1.grade)

	// // anonyms struct
	// var s1 = struct {
	// 	person
	// 	grade int
	// }{}
	// s1.person = person{"wick", 21}
	// s1.grade = 2

	// fmt.Println("name  :", s1.person.name)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)
	// fmt.Println("grade :", s1) // {{wick 21} 2}

	// // anonymous struct tanpa pengisian property
	// var s1 = struct {
	// 	person
	// 	grade int
	// }{}

	// // anonymous struct dengan pengisian property
	// var s2 = struct {
	// 	person
	// 	grade int
	// }{
	// 	person: person{"wick", 21},
	// 	grade:  2,
	// }

	var allStudent = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 23}, grade: 23},
		{person: person{"bond", 20}, grade: 2},
	}

	for _, each := range allStudent {
		fmt.Println(each)
	}
}

// type student struct {
// 	name  string
// 	grade int
// }

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}
