package main

import (
	"fmt"
	"sync"
)

func ByChan() {
	arr := []int{2, 4, 6, 8, 10}
	res := make(chan int)
	for i := range arr {
		go func() {
			res <- (arr[i] * arr[i])
		}()
		fmt.Println(<-res)
	}
}

func ByWg() {
	arr := []int{2, 4, 6, 8, 10}
	var res []int
	var wg sync.WaitGroup
	for i := range arr {
		wg.Add(1)
		num := arr[i]
		go func() {
			defer wg.Done()
			res = append(res, num*num)
		}()
	}
	wg.Wait()
	for i := range res {
		fmt.Println(res[i])
	}
}

func ByTwoChan() {
	arr := []int{2, 4, 6, 8, 10}
	in := make(chan int)
	out := make(chan int)
	go func() {
		for _, num := range arr {
			in <- num
		}
		close(in)
	}()
	go func() {
		for num := range in {
			out <- (num * num)
		}
		close(out)
	}()
	for num := range out {
		fmt.Println(num)
	}
}

func main() {
	ByWg()
	ByChan()
	ByTwoChan()
}
