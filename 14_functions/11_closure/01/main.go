package main

import "fmt"

func main() {
    x := 42
    fmt.Println(x)
    {
        fmt.Println(x)
        y := "This is in the inner scope"
        fmt.Println(y)
    }
    // Y is not accssiable from here
}