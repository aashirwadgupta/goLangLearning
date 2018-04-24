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
    title string
	name     string
	greeting string
}

func makeStaticMap() (map[string]Salutation){
    prefixStaticMap := map[string]Salutation {
        "1.":{"Mr.", "Aashirwad", "Hello!"},
	    "2.":{"Mr.", "Gupta", "Dear!"},
	    "3.":{"Sir.", "Aashirwad", "Hey"},
	    "4.":{"Mr.", "Aashirwad", "Hola!"},
    }
    
    return prefixStaticMap
}

func main() {
	var s = Salutation{"Mr.", "Aashirwad", "Hello!"}
	
	//Preparing a static map - map doesn't retains order
	fmt.Println(makeStaticMap())
	
	var statMap = makeStaticMap()
	//Updating the map by one more insertion
	statMap["5."]=s
	fmt.Println(statMap)
	
	//deleting the element from the map
	delete(statMap, "3.")
	fmt.Println(statMap)
}
