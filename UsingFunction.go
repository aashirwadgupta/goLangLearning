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
}

func main() {
	fmt.Println("Hello, playground")
	var s = Salutation{"Bob", "Hello!"}
	Greet(s)
}
