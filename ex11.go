package main

import "fmt"

func main() {
	b := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32)
	for i := range b {
		key := int(b[i]) / 10
		m[key] = append(m[key], b[i])
	}
	for i := range m {
		val, ok := m[i]
		if ok {
			for j := range val {
				fmt.Println(i*10, " ", val[j])
			}
		}
	}
}
