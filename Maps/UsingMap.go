/**
 * @author Aashirwad gupta
 *
 * Map - https://golang.org/ref/spec#Map_types
 */

package main

import (
	"fmt"
)

type Salutation struct {
	title    string
	name     string
	greeting string
}


func Greet(sal Salutation, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(sal.name)
		fmt.Println(i)
		fmt.Println(sal.greeting)
	}
}

func makeMap(sal []Salutation) map[string]string {
	//declare
	var prefixMap map[string]string
	//initialize
	prefixMap = make(map[string]string)
	//insert
	for i := 0; i < len(sal); i++ {
		prefixMap[sal[i].greeting] = sal[i].title + " " + sal[i].name
	}
	return prefixMap
}

func makeForEachMap(sal []Salutation) map[string]string {
	var prefixMap map[string]string
	//initialize
	prefixMap = make(map[string]string)
	for _, w := range sal {
		prefixMap[w.greeting] = w.title + " " + w.name
	}
	return prefixMap
}

func main() {
	//var s = Salutation{"Mr.", "Aashirwad", "Hello!"}
	//Greet(s, 5)

	salArr := []Salutation{
		{"Mr.", "Aashirwad", "Hello!"},
		{"Mr.", "Gupta", "Dear!"},
		{"Sir.", "Aashirwad", "Hey"},
		{"Mr.", "Aashirwad", "Hola!"},
	}
	//Prevents order
	fmt.Println(salArr)
	//Does not prevents order
	fmt.Println(makeMap(salArr))
	//For each loop - map doesn't retains order
	fmt.Println(makeForEachMap(salArr))
}
