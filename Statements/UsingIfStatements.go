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
func Greet(sal Salutation, do Printer, isFormal bool) {
    message, alternative := MessageWith2Return(sal.greeting, sal.name)
    if isFormal{
        do(message)
    }
	do(alternative)
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

func MessageWith2Return(param1, param2 string) (string, string){
    var par1 = param1 + " is first parameter"
    var par2 = param2 + " is second parameter"
    return par1, par2
}

/*
func PrintCustom(s string, custom string){
    fmt.Println(s+custom)
}
*/

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
	//This will give error due to type of arguments
	//This will require a lot of modifications for the following functions Print and PrintLine
	//Greet(s, PrintCustom)
	Greet(s, CreatePrintFunction(" Github "), true)
}
