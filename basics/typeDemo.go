package main

import "fmt"

type Func func([]byte) uint64

func hashFunc(data []byte) uint64 {
	total := uint64(0)
	for _, data := range data {
		total = total + uint64(data)
	}
	return total
}

func main() {
	fmt.Println("Hello World...")
	var hashFunc Func = hashFunc
	totalSum := hashFunc([]byte("Hello World"))
	fmt.Println(totalSum)
}
