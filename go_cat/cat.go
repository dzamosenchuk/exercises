package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
	pathfile := flag.String("f", "file", "f")
	
	flag.Parse()

	fmt.Println("Name of file is ", *pathfile)

	dat, err := os.ReadFile(*pathfile)
    check(err)
    fmt.Print(string(dat))

}