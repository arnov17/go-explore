package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var number = 23
	// var reflectvalue = reflect.ValueOf(number)

	// fmt.Println(reflectvalue)
	// fmt.Println("tipe vari", reflectvalue.Type())
	// fmt.Println("tipe vari", reflectvalue.Interface())

	// if reflectvalue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variable, ", reflectvalue.Int())
	// }

	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()

}

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
