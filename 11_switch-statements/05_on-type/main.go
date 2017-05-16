package main

import "fmt"

// Switches on type
// -- normally we swithc on value of variable
// -- go allow you to swithc on type of variable

type contract struct {
    greeting string
    name     string
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
    switch x.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    case contract:
        fmt.Println("contract")
    default:
        fmt.Println("unknown")
    }
}


func main() {
    SwitchOnType(7)
    SwitchOnType("McLead")
    var t = contract{"Good to see you,", "Tim"}
    SwitchOnType(t)
    SwitchOnType(t.greeting)
    SwitchOnType(t.name)
}