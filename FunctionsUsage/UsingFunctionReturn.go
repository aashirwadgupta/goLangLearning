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
	fmt.Println(CreateMessage(sal.greeting, sal.name))
}


func CreateMessage(param1, param2 string) string{
    var joinedStr = param1 + " Joined with " +param2
    return joinedStr
}

func main() {
	fmt.Println("Hello, playground")
	var s = Salutation{"Bob", "Hello!"}
	Greet(s)
}
