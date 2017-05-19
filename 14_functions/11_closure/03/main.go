package main

import "fmt"

func main() {
    x := 0
    incrament := func() int {
        x++
        return x
    }
    fmt.Println(incrament())
    fmt.Println(incrament())
}