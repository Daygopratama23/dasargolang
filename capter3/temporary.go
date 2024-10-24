package main

import "fmt"

func main() {
	var point = 8480.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	}

}
