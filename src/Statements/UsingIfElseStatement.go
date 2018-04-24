/**
 * @author Aashirwad gupta
 *
 * If Statements - https://golang.org/ref/spec#If_statements
 */

package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

//If-Else
func Greet(sal Salutation, do Printer, isFormal bool) {
    message, alternative := MessageWith2Return(sal.greeting, sal.name)
    if isFormal{
        do(message)
    } else {
        do(alternative)
    }
}


//function with no return types - denoted with empty paranthesis
type Printer func(string) ()


func MessageWith2Return(param1, param2 string) (string, string){
    var par1 = param1 + " is first parameter"
    var par2 = param2 + " is second parameter"
    return par1, par2
}

/*
This function is created to prevent the existing functionality of Printer function 
leaving us to modify very less
Now this is ClosureFunction

We have function which creates another function varying by a parameter
*/
func CreatePrintFunction(custom string) Printer {
    return func(s string){
        fmt.Println(custom +" : "+ s)
    }
}


func main() {
	var s = Salutation{"Aashirwad Gupta", "Hello!"}
  
	Greet(s, CreatePrintFunction(" Github "), true)
}
