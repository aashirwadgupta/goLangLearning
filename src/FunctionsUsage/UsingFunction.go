/**
 * @author Aashirwad Gupta
 *
 * Functions Usage and Syntaxes - https://golang.org/doc/codewalk/functions/
 */

package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

func Greet(sal Salutation) {
	fmt.Println(sal.greeting + " " +sal.name)
}

func main() {
	fmt.Println("Hello, playground")
	var s = Salutation{"Aashirwad Gupta", "Hello!"}
	Greet(s)
}
