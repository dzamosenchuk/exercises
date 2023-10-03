package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
     //   panic(e)
	 fmt.Println("Error ", e)
	 os.Exit(1)
    }
}
func main() {
	pathfile := flag.String("f", " ", "f")
	
	flag.Parse()
	if *pathfile == " " {
		fmt.Println("Use -f for define file name and path")
		os.Exit(1)
	}
	fmt.Println("Name of file is ", *pathfile)

	dat, err := os.ReadFile(*pathfile)
    check(err)
    fmt.Print(string(dat))

}