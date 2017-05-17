package main

import "fmt"

func main() {
    // Small and Large number
    var smlNum int
    var lrgNum int

    // Get the number inputs
    fmt.Println("Input a number")
    fmt.Scanln(&smlNum)
    fmt.Println("Input a number larger than the previous number")
    fmt.Scanln(&lrgNum)

    // Ensure that the small number is actually smaller
    if smlNum > lrgNum {
        fmt.Println("The small number is not smaller than the large number")
    } else {
        // Print out the remainder
        fmt.Println("The remainder of ", lrgNum, " divided by ", smlNum, " is ", lrgNum%smlNum)
    }
}