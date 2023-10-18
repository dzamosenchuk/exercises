package main

import (
	"fmt"
	"os"
	"testing"
)

// Testing fuction of counting lines in a file
func TestCountLines (t *testing.T) {
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

// Testing fuction of counting NON empty lines in a file
func TestCountLinesB (t *testing.T) {
	//Arange
	nameFile := "1_test.txt"
	d1 := []byte("hello\ngo\n\n")
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
	result := countLinesB(nameFile)
	defer os.Remove(nameFile)
	//Assert
	if result != expected {
		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
	}
}

// Testing read and print file
func TestFileRead(t *testing.T) {
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
	defer f.Close()
	//Act
	result := fileRead(nameFile)
	defer os.Remove(nameFile)
	
	//Assert
	if len(result) == 0 {
		t.Errorf("Incorrect result. File didn't read.")
	}

}