package main

import "fmt"

/*
	Just a simple little excercise to keep how loop through a go slice check
	if it is even or odd using modulus and print it using format and print
	from Go's standard library.
*/

func main() {

	// New int slice is created and initialized to the numSlice variable.
	numSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Loop through the numSlice slice and determine if it is even or odd
	// using modlus and the print the results.
	for i := 0; i < len(numSlice); i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even!!\n", numSlice[i])
		} else {
			fmt.Printf("%v is odd!!\n", numSlice[i])
		}
	}
}
