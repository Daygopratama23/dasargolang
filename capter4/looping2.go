package main

import "fmt"

func main() {
	var i = 0
	for {
		fmt.Println("hasil", i)
		i++
		if i == 5 {
			break
		}
	}
}
