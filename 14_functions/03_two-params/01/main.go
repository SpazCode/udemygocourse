package main

import "fmt"

func main() {
    greet("Stuart", "Sullivan")
    greet("Stepehen", "Sullivan")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}