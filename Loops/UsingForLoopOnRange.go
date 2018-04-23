/**
 * @author Aashirwad gupta
 *
 * For Loops - https://golang.org/ref/spec#For_statements
 */

package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

/*
Go through every item and populate the i with the index and s with the value
*/
func ForLoopForRange(sal []Salutation) {
    for i, s := range sal {
        fmt.Println(s.name +" "+ s.greeting)
        fmt.Println(i)
    }
}


func main() {
	
	slice := []Salutation{
	    {"Aashirwad", "Gupta"},
	    {"Hello", "Aash"},
	    {"Mr.", "Gupta"},
	}
	ForLoopForRange(slice)
}
