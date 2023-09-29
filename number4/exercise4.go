package main

import (
	"fmt"
	"math"
)

const s string = "const a const"
func main() {
	fmt.Println(s)

	const new = 1000000000

	const z = 3e20 / new
	fmt.Println(z)

	fmt.Println(int64(z))

	fmt.Println(math.Sin(new))

	fmt.Println()
}