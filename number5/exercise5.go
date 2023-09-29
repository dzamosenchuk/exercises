package main

import "fmt"

func main() {
	var i int = 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}
	
	fmt.Println()
	for j := 8; j <= 20; j++ {
		fmt.Println(j)
	}
	
	fmt.Println()
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("Done")
}