package main

import "fmt"

// braces always needed in if and for statements
func main() {
	if i := 23; i%2 == 2 {
		fmt.Println("control 1")
	} else if i >= 345 {
		fmt.Println("control 3")
	} else {
		fmt.Println("control 4")
	}
}
