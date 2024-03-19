package main

import "fmt"

func main() {
	num1 := 5
	num2 := 10
	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)
	num1, num2 = num2, num1
	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)
}
