package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		key := i
		val := i + 1
		go func() {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			m[key] = val
		}()
	}
	wg.Wait()
	for i := 0; i < 10000000; i++ {
		val, ok := m[i]
		if ok {
			fmt.Println(i, ":   ", val)
		}
	}
}
