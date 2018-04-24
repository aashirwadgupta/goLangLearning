/**
 * @author Aashirwad gupta
 *
 * Slice - https://golang.org/ref/spec#Slice_types
 */

package main

import (
	"fmt"
)

type Salutation struct {
	title    string
	name     string
	greeting string
}

func makeStaticSlice() []int {
	prefixStaticSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return prefixStaticSlice
}

func main() {

	var sliceOne []int

	//defining capacity as 3
	sliceOne = make([]int, 3)
	sliceOne[0] = 1
	sliceOne[1] = 2
	sliceOne[2] = 3

	fmt.Println(sliceOne)

	fmt.Println(makeStaticSlice())
}
