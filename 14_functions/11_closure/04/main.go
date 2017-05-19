package main

import "fmt"

func wrapper() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
} 

func main() {
    incrament := wrapper()
    fmt.Println(incrament())
    fmt.Println(incrament())
}