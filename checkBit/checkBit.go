package main

import "fmt"

func main() {
	fmt.Println(isBitSet(8, 4))
	fmt.Println(isBitSet(8, 3))
	fmt.Println(isBitSet(16, 5))
}

func isBitSet(num int, pos int) int {
	return (num >> (pos - 1)) & 1
}
