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



//Function with named return 
func FunctionWith2NamedReturn(param1, param2 string) (par1 string, par2 string){
    par1 = param1 + " is first parameter in named return"
    par2 = param2 + " is second parameter in named return"
    return 
}

func main() {
	fmt.Println("Hello, playground")
	var s = Salutation{"Bob", "Hello!"}
  
	message, alternate :=FunctionWith2NamedReturn(s.greeting, s.name)
	//A runtime exception will be thrown if variable is declared but not used
	fmt.Println(message)
	fmt.Println(alternate)
  
  //If this has to be ignored then just use underscore like this below
  _, alternate2 :=FunctionWith2NamedReturn(s.greeting, s.name)
  fmt.Println(alternate2)
}
