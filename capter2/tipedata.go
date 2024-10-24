package main

import "fmt"

func main() {
	var bilanganA uint = 1 + 2
	var bilanganB = 2 + bilanganA
	fmt.Println("hasil", bilanganA, bilanganB)
	fmt.Print("john", "week\n")
	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
	// fmt.Printf("jhon%","week%")
}
