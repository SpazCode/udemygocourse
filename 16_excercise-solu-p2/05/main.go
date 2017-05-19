package main

import "fmt"

func main() {
    foo := func(nums ...int) {
        fmt.Println(len(nums))
    }

    foo(1, 2)
    foo(1, 2, 3)
    aSlice := []int{1, 2, 3, 4}
    foo(aSlice...)
    foo()
}