package main

import (
    "fmt"
    "log"
    "github.com/XihuanYang/greetings"
    "rsc.io/quote"
)

func main() {
    //importing packages from a published module,
    //quote in https://pkg.go.dev/
    fmt.Println(quote.Hello())
    
    //Use local package
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
   
    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}
 
    message, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
