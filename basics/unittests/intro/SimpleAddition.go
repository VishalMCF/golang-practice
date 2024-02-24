package main

import "fmt"

func main() {
	a := 7
	b := 10
	sumOfAandB := AddNumbers(a, b)
	fmt.Println(sumOfAandB)
}

func AddNumbers(a, b int) int {
	return a + b
}
