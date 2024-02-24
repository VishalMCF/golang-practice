package main

import "testing"

// single testcase in one method
func TestAddNumbers(t *testing.T) {
	input1, input2 := 3, 4
	expected := 7
	output := AddNumbers(input1, input2)
	if output != expected {
		t.Errorf("Expected '%d', but got '%d'", expected, output)
	}
}

// multiple test cases in one method
func TestMultipleInputsAddNumbers(t *testing.T) {
	type Input struct {
		number1, number2, expected int
	}

	inputs := []Input{
		{1, 2, 3},
		{2, 3, 5},
		{-1, -19, -20},
	}

	for _, input := range inputs {
		got := AddNumbers(input.number1, input.number2)
		if got != input.expected {
			t.Errorf("Expected '%d', but got '%d'", input.expected, got)
		}
	}
}
