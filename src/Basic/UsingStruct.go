/**
 * @author Aashirwad Gupta
 *
 * Struct and it's syntax - https://golang.org/ref/spec#Struct_types
 */

package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

func main() {
	fmt.Println("Hello, playground")
	var sal = Salutation{"Bob", "Hello!"}
  fmt.Println(sal.greeting + " " +sal.name)
	
}
