package main

import (
	"fmt"
	"math"
)

func main() {
	const a = "abc"
	fmt.Println(a)

	// integer constants are stored as aribitrary types like 2e2
	const b = 245e4
	fmt.Println(b)

	const c int = 30948587
	fmt.Println(c)

	fmt.Println(math.Sin(b))
}
