package main

import "fmt"

func main() {
	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Sup my friend with 2 names")
	case myFriendsName == "Tim":
		fmt.Println("Sup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Sup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Sup Marcus... or er, Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Sup Julian")
		fallthrough
	case myFriendsName == "Sushant":
		fmt.Println("Sup Sushant")
	default:
		fmt.Println("Nothing matched; this is the default")
	}	
}