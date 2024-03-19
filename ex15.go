package main

import (
	"fmt"
	"reflect"
)

func GetType(vr interface{}) reflect.Type {
	return reflect.TypeOf(vr)
}

func main() {
	var a float64
	var b int
	var c string
	var g map[int]int
	fmt.Println(GetType(a))
	fmt.Println(GetType(b))
	fmt.Println(GetType(c))
	fmt.Println(GetType(g))
}
