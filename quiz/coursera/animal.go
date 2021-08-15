package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	colection := map[string]Animal{
		"cow":   Animal{"grass", "walk", "mooo"},
		"bird":  Animal{"worms", "fly", "peepee"},
		"snake": Animal{"mice", "slither", "pspsps"},
	}
	for {
		fmt.Print(">")
		animal := "0"
		action := "0"
		fmt.Scan(&animal, &action)
		if action == "eat" {
			colection[animal].Eat()
		} else if action == "move" {
			colection[animal].Move()
		} else if action == "speak" {
			colection[animal].Speak()
		}
	}
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	animals := map[string]Animal {
		"cow": {food: "grass", locomotion: "walk", noise: "moo"},
		"bird": {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		input := strings.TrimSpace(strings.ToLower(scanner.Text()))
		values := strings.Fields(input)
		if len(values) == 2 {
			animal, ok := animals[values[0]]
			if ok {
				switch info := values[1]; info {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Printf("info %s is not a valid request\n", info)
				}
			} else {
				fmt.Printf("Animal %s does not exist\n", values[0])
			}
		} else {
			fmt.Println("Invalid request. Specify <animal> <info>.")
		}
	}
}

*/
