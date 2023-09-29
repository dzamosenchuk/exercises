package main

import "fmt"

func main () {
// #1 hello world 
	fmt.Println("Hello World")
	fmt.Println()


// #2 Values

	fmt.Println("go" + "lang")
	fmt.Println("2+2=", 1+1)
	fmt.Println("5.0/8.0=", 5.0/8.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println()

// #3 Variables

	var first = "string first"
	fmt.Println(first)

	var a, b int = 3, 56
	fmt.Println(a, b)

	var lie = true
	fmt.Println(lie)

	var z int
	fmt.Println(z)

	new := "new variable"
	fmt.Println(new)

	fmt.Println()

}