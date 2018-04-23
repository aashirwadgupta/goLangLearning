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
