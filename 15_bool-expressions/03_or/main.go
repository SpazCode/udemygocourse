package main

import "fmt"

func main() {
    if true || false {
        fmt.Println("This ran")
    }

    x := 14

    if x == 14 || false {
        fmt.Println("This ran too")
    }
}