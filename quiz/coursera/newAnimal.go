package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func scanUserInput(prompt string) (string, error) {
	// Prompt user to enter input
	fmt.Print(prompt)

	// Read the user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Error handling
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func main() {

	var newAnimals = make(map[string]Animal)

	for {
		// Enter command
		input, err := scanUserInput("> ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		values := strings.Fields(input)

		// Check inputs
		if len(values) != 3 {
			fmt.Println("Enter 'command param1 param2'")
			continue
		}

		// Process command
		if values[0] == "newanimal" {
			fmt.Println("command:", values[0])
			fmt.Println("name of new animal:", values[1])
			fmt.Println("type of the new animal:", values[2])
			fmt.Println("----------------------------------------")

			switch values[2] {
			case "cow":
				newAnimals[values[1]] = Cow{}
			case "bird":
				newAnimals[values[1]] = Bird{}
			case "snake":
				newAnimals[values[1]] = Snake{}
			default:
				fmt.Printf("Unknown animal type: %s\n", values[2])
				continue
			}

			fmt.Println("Created it!")

		} else if values[0] == "query" {
			fmt.Println("command:", values[0])
			fmt.Println("name of the animal:", values[1])
			fmt.Println("information of the animal:", values[2])
			fmt.Println("----------------------------------------")

			if animal, ok := newAnimals[values[1]]; ok {
				// Check info
				switch values[2] {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Printf("Unknown request: %s\n", values[2])
				}
			} else {
				fmt.Printf("animal '%s' not found\n", values[1])
			}

		} else {
			fmt.Println("Unknown command. Supported command: 'newanimal' 'query'")
		}
	}
}

/*
package main

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (a animal) Eat() {
	fmt.Println(a.food)
	return
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
	return
}

func (a animal) Speak() {
	fmt.Println(a.noise)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}

	var generalAnimal animalInterface
	for {

		var command, reqAnimal, reqType string
		fmt.Print(">")
		fmt.Scan(&command, &reqAnimal, &reqType)

		if command == "query" {
			generalAnimal = animalMap[reqAnimal]
			switch reqType {
			case "eat":
				generalAnimal.Eat()
			case "move":
				generalAnimal.Move()
			case "speak":
				generalAnimal.Speak()
			default:
				fmt.Println("Unknown request type")
			}
		}

		if command == "newanimal" {
			animalMap[reqAnimal] = animalMap[reqType]
			fmt.Println("Created it!")
		}
	}
}

*/
