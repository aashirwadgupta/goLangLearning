package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}



//Variadic Function 
func VariadicFunction(param1 string, param2 ...string) {
    
    fmt.Println(len(param2))
    
    fmt.Println(param1 + " is first parameter in VariadicFunction")
    fmt.Println(param2[0] + " is zeroth parameter in VariadicFunction's Slice which is similar to array")
    fmt.Println(param2[1] + " is first parameter in VariadicFunction's Slice which is similar to array")
    fmt.Println(param2[2] + " is second parameter in VariadicFunction's Slice which is similar to array")
     
}

func main() {
	fmt.Println("Hello, playground")
	var s = Salutation{"Aashirwad", "Hello!"}
  
	VariadicFunction(s.greeting, s.name, "Aash", "Gupta")
}
