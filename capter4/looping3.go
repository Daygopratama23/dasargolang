package main

import "fmt"

func main() {
	// var i = 0
	var xs = "01234"
	for i, v := range xs {

		fmt.Println("index", i, "value=", v)

		var ys = xs
		for _, v := range ys {
			fmt.Println("value=", v, "value", ys, i)
		}

	}
}
