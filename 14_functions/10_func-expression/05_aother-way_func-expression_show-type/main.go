package main

import "fmt"


func makeGreeter() func() {
    return func() {
        fmt.Println("Hellow World!")
    }
}

func main() {
    greeting := makeGreeter()
    greeting()
    fmt.Printf("%T\n", greeting)
}