/**
 * @author Aashirwad gupta
 *
 * For Loops - https://golang.org/ref/spec#For_statements
 */

package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}


func ForLoop(sal Salutation, times int) {
    for i :=0; i<times; i++ {
        fmt.Println(sal.name)
        fmt.Println(i)
        fmt.Println(sal.greeting)
    }
}

func ForLikeWhile(sal Salutation, times int) {
    i := 0
    for  i<times{
        fmt.Println(i)
        fmt.Println(sal.name)
        fmt.Println(sal.greeting)
        i++ 
    }
}

/*
func ForWithoutCondition(sal Salutation, times int) {
   
    for {
        fmt.Println(i)
        fmt.Println(sal.name)
        fmt.Println(sal.greeting)
        i++ 
    }
}
*/

func main() {
	var s = Salutation{"Aashirwad", "Gupta"}
	//ForLoop(s, 5)
	ForLikeWhile(s, 2)
}
