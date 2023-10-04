package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

//check errors
func check(e error) { 
	if e != nil {
		fmt.Println("Error is ", e)
		os.Exit(1)
	}
}

//count all lines in file
func countLines(fileName string) { 
	file, err := os.Open(fileName)
	check(err)
	countLines := 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		countLines++
	}
	fmt.Println("Number of lines is ", countLines)
}

//count not empty lines in file
func countLinesB(fileName string) { 
	file, err := os.Open(fileName)
	check(err)
	countLines := 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			countLines++
		}

	}
	fmt.Println("Number of lines is ", countLines)
}

//read and print file
func fileRead(fileName string) { 
	dat, err := os.ReadFile(fileName)
	check(err)
	fmt.Print(string(dat))
}

func main() {
	pathfile := flag.String("f", "", "f")
	countStrings := flag.String("n", "", "n")
	countStringsB := flag.String("b", "", "b")

	flag.Parse()

	switch {
	case len(*pathfile) > 0:
		fileRead(*pathfile)

	case len(*countStrings) > 0:
		countLines(*countStrings)

	case len(*countStringsB) > 0:
		countLinesB(*countStringsB)

	default:
		fmt.Println("Any flags are defined. Use -f for print a file, -n for count strings in a file, -b for count all non empty strings.")
		os.Exit(1)
	}

}
