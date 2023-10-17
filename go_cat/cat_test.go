package main

import (
	"testing"
)
// Testing fuction of counting lines in a file
func TestCountLines (t *testing.T) {
	//Arange
	nameFile := "1_test.txt"
	expected := 5

	//Act
	result := countLines(nameFile)
	
	//Assert
	if result != expected {
		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
	}
}

// Testing fuction of counting NON empty lines in a file
func TestCountLinesB (t *testing.T) {
	//Arange
	nameFile := "1_test.txt"
	expected := 5

	//Act
	result := countLinesB(nameFile)
	
	//Assert
	if result != expected {
		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
	}
}

// Testing read and print file
func TestFileRead(t *testing.T) {
	//Arange
	nameFile := "1_test.txt"

	//Act
	result := fileRead(nameFile)
	//Assert
	if len(result) == 0 {
		t.Errorf("Incorrect result. File didn't read.")
	}

}