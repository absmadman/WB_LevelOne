package main

import "fmt"

func AppendMap(n []int) *map[int]int {
	m := make(map[int]int)
	for i := range n {
		m[n[i]] = n[i]
	}
	return &m
}

func main() {
	n1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := []int{5, 6, 7, 8}
	n3 := []int{}
	m1 := *AppendMap(n1)
	m2 := *AppendMap(n2)
	for i := range m1 {
		if _, ok := m2[i]; ok {
			n3 = append(n3, i)
		}
	}
	for i := range n3 {
		fmt.Println(n3[i])
	}
}
