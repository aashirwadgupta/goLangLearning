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

func Greet(sal Salutation, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(sal.name)
		fmt.Println(i)
		fmt.Println(sal.greeting)
	}
}

func main() {
	var s = Salutation{"Aashirwad", "Hello!"}
	Greet(s, 5)
}
