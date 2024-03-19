package main

import (
	"fmt"
	"os"
)

func SetToOne(numbit int, number int) int {
	mask := 1 << numbit
	return number | mask
}

func SetToZero(numbit int, number int) int {
	mask := 1 << numbit
	return ^(^number | mask)
}

func main() {
	number := 128128128
	bit := 0
	fmt.Fscan(os.Stdin, &bit)
	fmt.Println(SetToZero(bit, number))
	fmt.Println(SetToOne(bit, number))
}
