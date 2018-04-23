/**
 * @author Aashirwad gupta
 *
 * Switch Statements - https://golang.org/ref/spec#Switch_statements
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

*/
func Greet(sal Salutation, do Printer, isFormal bool) {
    message, alternative := MessageWith2Return( sal.name, sal.greeting)
    if isFormal{
        //If true pass in the name
        //prefix := GetPrefixViaNothing(sal.name);
        prefix := GetPrefix(sal.name);
        do(prefix + message)
    } else {
        //If false pass greeting as name
        //prefix := GetPrefixViaNothing(sal.name);
        prefix := GetPrefix(sal.greeting);
        do(prefix + alternative)
    }
    
    
}

/*
Using fallthrough switches and executes the immediate next
Also we can have multiple checks of the type
For example - case "Aashirwad", "Gupta", "Hello!"
*/
func GetPrefix(name string) (prefix string) {
    switch name{
        case "Aashirwad", "Hello!": prefix = "Mr. "
        case "Aash": prefix = "Buddy "
        default: prefix = "Sir. "
    }
    return
}

/*
We can also use a switch case with no initial condition. 
The default treatment is true so that case are evaluated
now in the case we can include condition which may evaluate to true or false
*/
func GetPrefixViaNothing(name string) (prefix string) {
    switch {
        case len(name)>5, name=="Hello!", name=="Aashirwad": prefix = "Mr. "
        case name=="Aash", len(name)<=5: prefix = "Buddy "
        default: prefix = "Sir. "
    }
    return
}

//function with no return types - denoted with empty paranthesis
type Printer func(string) ()


func MessageWith2Return(param1, param2 string) (string, string){
    var par1 = param1 + " is in if condition"
    var par2 = param2 + " is in else condition"
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
	var s = Salutation{"Aashirwad", "Hello!"}
  
	Greet(s, CreatePrintFunction(" Github "), true)
	
	var s2 = Salutation{"Hello!", "Aash"}
  
	Greet(s2, CreatePrintFunction(" Github "), false)
	
	var s3 = Salutation{"Gupta", "Aashirwad"}
  
	Greet(s3, CreatePrintFunction(" Github "), true)
}
