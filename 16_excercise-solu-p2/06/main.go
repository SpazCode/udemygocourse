package main

import "fmt"

func main() {
    fmt.Println(scan(1, 2, 3, 4, 5, 6, 7, 8, 9, 0))
}

func scan(series ... int) int {
    var maxProduct int
    if len(series) < 4 {
        return 0
    }
    for i := 3; i < len(series); i++ {
        // fmt.Println(series[i] * series[i-1] * series[i-2] * series[i-3])
        if maxProduct < (series[i] * series[i-1] * series[i-2] * series[i-3]) {
            maxProduct = (series[i] * series[i-1] * series[i-2] * series[i-3])
        }
    }
    return maxProduct
}

// Largest product in a series