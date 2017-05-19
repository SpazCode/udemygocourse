package main

import "fmt"

const p string = "death & taxes"
const q string = "war & peace"

func main() {
	// const p = 30
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}