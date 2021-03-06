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

//func is a part of type definition of the function do
//Stage 2
func Greet(sal Salutation, do Printer) {
	do(sal.greeting)
	do(sal.name)
}

//func is a part of type definition of the function do
/* Stage 1
func Greet(sal Salutation, do func(string) {
	do(sal.greeting)
	do(sal.name)
}
*/

//function with no return types - denoted with empty paranthesis
type Printer func(string) ()

func Print(s string){
    //Since below methods are a bit different from the signature of the the typed method,
    //we can't use them directly
    fmt.Print("From Print Function "+s)
}

func PrintLine(s string){
    //Since below methods are a bit different from the signature of the the typed method,
    //we can't use them directly
    fmt.Println("From PrintLine Function "+s)
}

func main() {
	var s = Salutation{"Aashirwad Gupta", "Hello!"}
	Greet(s, Print)
	fmt.Println(" Now calling line function ")
	Greet(s, PrintLine)
}
