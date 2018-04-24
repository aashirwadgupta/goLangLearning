package greeting

import (
    "fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Salutations []Salutation

func (sal Salutations) Greet(do Printer, isFormal bool, times int) {
    for _, s := range sal{
        message, alternate := CreateMessage(s.Name, s.Greeting)
        if prefix := GetPrefix(s.Name); isFormal {
            do(prefix+message)
        } else {
            do(alternate)
        }
    }
	
}

func GetPrefix(name string) (prefix string) {
    prefixMap := map[string]string{
        "Aash":"Dear. ",
        "Aashirwad":"Mr. ",
        "Gupta":"Sir. ",
    }
    return prefixMap[name]
}

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

func CreatePrintFunction(custom string) Printer {
    return func(s string){
        fmt.Println(s+custom)
    }
}

func CreateMessage(param1, param2 string) (string, string){
    var par1 = param1 + " is first parameter"
    var par2 = param2 + " is second parameter"
    return par1, par2
}