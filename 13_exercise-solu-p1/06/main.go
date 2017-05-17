package main

import "fmt"

func main() {
    for i := 0; i <= 100; i++ {
        switch {
            case i%3 + i%5 == 0: 
                fmt.Print("FizzBuzz")
            case i%3 == 0: 
                fmt.Print("Fizz")
            case i%5 == 0:
                fmt.Print("Buzz")
            default:
                fmt.Print(i)
        }
        fmt.Print("\n")
    }
}