package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// check errors
func check(e error) {
	if e != nil {
		fmt.Println("Error is ", e)
		os.Exit(1)
	}
}

// Count bytes in a file
func countBytes(fileName string) int64 {
	countBytes, err := os.Stat(fileName)
	check(err)
	return countBytes.Size()
}

// count all lines in file
func countLines(fileName string) int {
	file, err := os.Open(fileName)
	check(err)
	countLines := 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		countLines++
	}
	return countLines
}

func main() {

	cFlag := flag.String("c", "", "-c option")
	lFlag := flag.String("l", "", "-l option")

	flag.Parse()

	switch {
	case len(*cFlag) > 0:
		fmt.Println("Number of bytes is ", countBytes(*cFlag))
	case len(*lFlag) > 0:
		fmt.Println("Number of lines is ", countLines(*lFlag))
	default:
		fmt.Println("Any flags are defined. Use -c count number of bytes a file, -l count number of lines in a file.")
		os.Exit(1)
	}
}
