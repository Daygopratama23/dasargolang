package main

import "fmt"

func main() {
	var point = 1

	switch point {
	case 8:
		fmt.Println("angka 8")
	case 2, 3, 4, 5:
		fmt.Println("angka bukan 8")
	}

}
