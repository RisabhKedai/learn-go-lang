package main

import "fmt"

func main() {
	// normal for loops
	i := 1
	for i <= 2 {
		fmt.Println("Test", i)
		i += 1
	}
	for i := 1; i < 2; i += 1 {
		fmt.Println("iter ", i)
	}

	// range runs from 0 to the number on RHS
	for i := range 10 {
		fmt.Println("iter", i)
	}

	// infinite loop
	for {
		fmt.Println("test", i)
		break
	}
}
