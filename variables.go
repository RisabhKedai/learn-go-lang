package main

import "fmt"

func main() {
	var a = "sample string"
	fmt.Println(a)

	var b, c = 1, "peitigj"
	fmt.Println(b, c)

	// short hand for declaring variables
	// only available inside function bodies
	d := 2.4
	fmt.Println(d)

	// declaration with type
	var e int = 343
	fmt.Println(e)

	// failure scenario
	var f int = 56
	fmt.Println(f)

	var g string = "433"
	fmt.Println(g)

	var h string
	fmt.Println(h)
}
