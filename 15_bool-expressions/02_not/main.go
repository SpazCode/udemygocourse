package main

import "fmt"

func main() {
    if !true {
        fmt.Println("This won't Ran")
    }

    if !false {
        fmt.Println("This Ran")
    }
}