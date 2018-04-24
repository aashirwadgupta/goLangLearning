/**
 * @author Aashirwad gupta
 *
 * Slice - https://golang.org/ref/spec#Slice_types
 */

package main

import (
	"Methods/greeting"
)


func makeStaticSlice() ([]int){
    prefixStaticSlice := []int{1,2,3,4,5,6,7,8,9,10}
    return prefixStaticSlice
}

func main() {
	
	

	sliceOne := greeting.Salutations{
	    {"Aashirwad", "Gupta"},
	    {"Gupta", "Aashirwad"},
	    {"Aash", "Gupta"},
	}
	sliceOne.Greet(greeting.CreatePrintFunction("Github"), true, 6)
}
