package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	type Name struct {
		fname string
		lname string
	}
	var filename string
	var nameSlice []Name

	nameSlice = make([]Name, 0)
	fmt.Println("Give path of file to read: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entryName := new(Name)
		fullName := scanner.Text()
		splitstr := strings.SplitN(fullName, " ", 2)

		if len(splitstr[0]) > 20 {
			splitstr[0] = splitstr[0][:20]
		}
		if len(splitstr[1]) > 20 {
			splitstr[1] = splitstr[1][:20]
		}

		entryName.fname = splitstr[0]
		entryName.lname = splitstr[1]
		nameSlice = append(nameSlice, *entryName)
	}

	for i := 0; i < len(nameSlice); i++ {
		fmt.Println("First name: ", nameSlice[i].fname, "   Last name: ", nameSlice[i].lname)
	}

}

func check(error error) {
	if error != nil {
		panic(error)
	}
}
