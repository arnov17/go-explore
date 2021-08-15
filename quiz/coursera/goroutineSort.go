package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortSlice(slice []int, wg *sync.WaitGroup, goRoutineId int) []int {
	defer wg.Done()
	sort.Ints(slice)
	fmt.Printf("Go Routine %d ; I am sorting %v \n", goRoutineId, slice)
	fmt.Println("-------------")
	return slice
}

func scanUserInput(prompt string) (string, error) {
	fmt.Println("-------------")
	fmt.Println(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func main() {
	// User input
	input, err := scanUserInput("Enter integers (separate each number with space. (18 17 16 15 99 14 13 12 11 10 9 8 7 6 5 4 54):")
	fmt.Println("------------------")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringSlice := strings.Split(input, " ")
	slice := make([]int, len(stringSlice))
	for i := range slice {
		n, err := strconv.Atoi(strings.TrimSpace(stringSlice[i]))
		if err != nil {
			panic("input format with spaces")
		}
		slice[i] = n
	}
	sliceCount := len(slice)
	partition := 4
	size := sliceCount / partition

	fmt.Println(slice)

	wg := sync.WaitGroup{}
	wg.Add(partition)

	go func() {
		for start := 0; start < sliceCount; start += size {
			end := start + size
			if end > len(slice) {
				break
			}
			if end+size > sliceCount {
				end = sliceCount
			}
			go sortSlice(slice[start:end], &wg, start)
		}
	}()
	wg.Wait()
	sort.Ints(slice)
	fmt.Println("\nSorted Output : ")
	fmt.Println(slice)
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortArr(arr []int, c *chan []int, goRoutineId int) {
	fmt.Printf("Go Routine %d ; I am sorting %v \n", goRoutineId, arr)
	sort.Ints(arr)
	//fmt.Printf("Go Routine %d ; I am sending %v \n", goRoutineId, arr)
	*c <- arr
}

func main() {

	fmt.Print("Enter space seperate list of integers ")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	n := string(line)
	fmt.Println("")
	fmt.Println("")
	//n := "21 20 19 18 17 16 15 99 14 13 12 11 10 9 8 7 6 5 4 3 2 1"
	var arr []int

	narr := strings.Split(strings.TrimSpace(n), " ")

	for _, v := range narr {
		val, _ := strconv.ParseInt(v, 10, 64)
		arr = append(arr,int(val) )
	}

	div := len(arr) / 4
	c := make(chan []int)

	goRoutineId := 0
	go sortArr(arr[:div], &c, goRoutineId)
	goRoutineId++
	go sortArr(arr[div:2*div], &c, goRoutineId)
	goRoutineId++
	go sortArr(arr[2*div:3*div], &c, goRoutineId)
	goRoutineId++
	go sortArr(arr[3*div:], &c, goRoutineId)

	sortedArr := merge(merge(<-c, <-c), merge(<-c, <-c))

	fmt.Println("\nSorted Output : ")
	fmt.Println(sortedArr)

}

func merge(arr1 []int, arr2 []int) []int {
	var result []int
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}
	for len(arr1) > 0 {
		result = append(result, arr1[0])
		arr1 = arr1[1:]
	}
	for len(arr2) > 0 {
		result = append(result, arr2[0])
		arr2 = arr2[1:]
	}
	return result
}


*/

/*
koji

// Write a program to sort an array of integers.
// The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
// When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calcChunkSize(length, div int) []int {

	chunkSize := make([]int, 0)

	for i := 0; i < length%div; i++ {
		chunkSize = append(chunkSize, length/div+1)
	}

	for i := 0; i < (div - length%div); i++ {
		chunkSize = append(chunkSize, length/div)
	}

	return chunkSize
}

func divide(slice, chunkSize []int) [][]int {

	chunks := make([][]int, 0)

	i := 0
	for {
		if len(slice) == 0 {
			break
		}

		chunks = append(chunks, slice[0:chunkSize[i]])
		slice = slice[chunkSize[i]:]
		i++
	}

	return chunks
}

func sortSubarray(a []int, c chan []int) {

	fmt.Println("Sorting subarray:", a)

	tmp := make([]int, len(a))
	copy(tmp, a)

	sort.Ints(a)

	// fmt.Println("Subarray (before sorting -> sorted):", tmp, "->", a)

	c <- a // send sorted array to c
}

func merge(a, b []int) []int {

	merged := make([]int, 0)

	merged = append(merged, a...)
	merged = append(merged, b...)

	sort.Ints(merged)

	return merged
}

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

func main() {

	numbers := make([]int, 0)
	c := make(chan []int)

	// User input
	input, err := scanUserInput("Enter integers (separate each number with space. e.g. 100 -2 31 46):")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := strings.Fields(input)

	// DEBUG: Display entered values
	for i := 0; i < len(values); i++ {
		value, err := strconv.Atoi(values[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		numbers = append(numbers, value)
	}
	// fmt.Println("Entered integers:", numbers)

	// Devide input array to 4 arrays
	chunkSize := calcChunkSize(len(numbers), 4)
	divided := divide(numbers, chunkSize)
	fmt.Println("\nSubarrays that will be sorted:")
	fmt.Println(divided)
	fmt.Println()

	// Start goroutines to sort subarray
	for i := 0; i < 4; i++ {
		go sortSubarray(divided[i], c)
	}

	// Receive sorted subarray
	var sortedSubarray [4][]int
	for i := 0; i < 4; i++ {
		sortedSubarray[i] = <-c
	}
	fmt.Println("\nSorted subarrays:")
	fmt.Println(sortedSubarray)
	fmt.Println()

	// Merge
	result := merge(merge(sortedSubarray[0], sortedSubarray[1]), merge(sortedSubarray[2], sortedSubarray[3]))
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println()
}

*/
