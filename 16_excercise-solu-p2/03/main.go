package main

import "fmt"

func maximum(numbers ...int) int {
    var max int
    for _, n := range numbers {
        if n > max {
            max = n
        }
    }
    return max
}

func main() {
    fmt.Println(maximum(10, 3, 15, 6, 7, 20, 18))
}