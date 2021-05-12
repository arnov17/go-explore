package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Arnov Julian", "ov"))
	fmt.Println(strings.Contains("Arnov Julian", "Ani"))
	fmt.Println(strings.Split("Arnov Julian", " "))
	fmt.Println(strings.ToLower("Arnov Julian"))
	fmt.Println(strings.ToUpper("arnov julian"))
	fmt.Println(strings.Trim("     Arnov Julian   ", " "))
	fmt.Println(strings.ReplaceAll("ani nia ina", "ani", "anna"))

}
