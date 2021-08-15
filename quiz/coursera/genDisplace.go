package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, vel, disp float64) func(float64) float64 {

	fn := func(t float64) float64 {

		return acc*math.Pow(t, 2)/2 + vel*t + disp
	}
	return fn
}

func main() {
	fmt.Println("Give initial values for acceleration, velocity and displacement")

	data := [3]float64{}
	for i := 0; i < 3; i++ {
		fmt.Scan(&data[i])
	}
	fn := GenDisplaceFn(data[0], data[1], data[2])
	fmt.Println("Give the time for the calculation")
	var t float64
	fmt.Scan(&t)

	fmt.Printf("The displacement after time %.3f is %.3f \n", t, fn(t))

}

/*
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func scanUserInput(prompt string) (string, error) {
	// Prompt user to enter input
	fmt.Println(prompt)

	// Read the user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Error handling
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		// s = ½ a t2 + vot + so
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {

	// Enter values for acceleration, initial velocity, and initial displacement.
	input, err := scanUserInput("Enter acceleration initialVelocity initialDisplacement:\n(Separate with spaces. e.g. 10 2 1)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := strings.Fields(input)

	// DEBUG: Display entered values
	fmt.Println("acceleration:", values[0])
	fmt.Println("initialVelocity:", values[1])
	fmt.Println("initialDisplacement:", values[2])

	// Parse values
	acceleration, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	initialVelocity, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	initialDisplacement, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	// Get a displacement function with the initial values
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// Compute displacement for entered time value
	for {
		input, err = scanUserInput("Enter time (Enter X to exit):")
		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}

		// Check if 'X' is entered
		inputRune := []rune(input)
		if len(input) == 1 && inputRune[0] == 'X' {
			fmt.Println("User entered X. Exitting...")
			break
		}

		t, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(6)
		}

		fmt.Printf("Displacement after %f seconds:\n", t)
		fmt.Println(fn(t))
		fmt.Println()
	}
}

*/
