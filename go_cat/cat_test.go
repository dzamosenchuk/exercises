package main

import (
	"testing"
)
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
