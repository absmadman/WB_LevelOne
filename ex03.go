package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func ByChan() {
	arr := []int{2, 4, 6, 8, 10}
	res := make(chan int)
	sum := 0
	for i := range arr {
		go func() {
			res <- (arr[i] * arr[i])
		}()
		sum += <-res
	}
	fmt.Println(sum)
}

func SumByWg() {
	arr := []int{2, 4, 6, 8, 10}
	var res []int
	sum := 0
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
		sum += res[i]
	}
	fmt.Println(sum)
}

func ByTwoChan() {
	arr := []int{2, 4, 6, 8, 10}
	in := make(chan int)
	out := make(chan int)
	sum := 0
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
		sum += num
		//fmt.Println(num)
	}
	fmt.Println(sum)
}

func ByMutex() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := range arr {
		wg.Add(1)
		num := arr[i]
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			sum += num * num
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}

func ByAtomic() {
	arr := []int{2, 4, 6, 8, 10}
	var sum int32
	var wg sync.WaitGroup
	for i := range arr {
		wg.Add(1)
		num := arr[i]
		go func() {
			defer wg.Done()
			atomic.AddInt32(&sum, int32(num*num))
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}

func main() {
	SumByWg()
	ByChan()
	ByTwoChan()
	ByMutex()
	ByAtomic()
}
