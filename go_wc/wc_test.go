package main

import (
	"fmt"
	"os"
	"testing"
)

// Testing fuction of counting lines in a file
func TestCountLines(t *testing.T) {
	//Arange
	nameFile := "1_test.txt"
	d1 := []byte("hello\ngo\n")
	f, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(1)
	}

	err = os.WriteFile(nameFile, d1, 0644)

	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(1)
	}

	expected := 2
	defer f.Close()
	//Act
	result := countLines(nameFile)
	defer os.Remove(nameFile)
	//Assert
	if result != expected {
		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	//Arange
	nameFile := "1_test.txt"
	d1 := []byte("hello\ngo\n")
	f, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(1)
	}

	err = os.WriteFile(nameFile, d1, 0644)

	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(1)
	}

	var expected int64 = 9
	defer f.Close()
	//Act
	result := countBytes(nameFile)
	defer os.Remove(nameFile)
	//Assert
	if result != expected {
		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
	}
}
