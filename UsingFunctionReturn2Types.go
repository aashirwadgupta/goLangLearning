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

func MessageWith2Return(param1, param2 string) (string, string){
    var par1 = param1 + " is first parameter"
    var par2 = param2 + " is second parameter"
    return par1, par2
}

func main() {
	fmt.Println("Hello, playground")
	var s = Salutation{"Bob", "Hello!"}
	Greet(s)
	message, alternate :=MessageWith2Return(s.greeting, s.name)
	//A runtime exception will be thrown if variable is declared but not used
	fmt.Println(message)
	fmt.Println(alternate)
  
  //If this has to be ignored then just use underscore like this below
  _, alternate2 :=MessageWith2Return(s.greeting, s.name)
  fmt.Println(alternate2)
}
