package main

import (
	"context"
	"fmt"
	"time"
)

type Worker struct {
	id   int
	jobs <-chan int
	res  chan<- int
}

func NewWorker(id int, jobs <-chan int, res chan<- int) *Worker {
	return &Worker{
		id:   id,
		jobs: jobs,
		res:  res,
	}
}

func (w *Worker) Work() {
	for num := range w.jobs {
		w.res <- (num * num)
	}
}

func ByChannels() {
	num := 20
	jnum := 100
	jobs := make(chan int, jnum)
	res := make(chan int, jnum)
	flag := make(chan int)
	defer close(res)
	defer close(jobs)
	for i := 1; i < num; i++ {
		worker := NewWorker(i, jobs, res)
		go worker.Work()
	}
	go func() {
		for i := 0; ; i += 1 {
			go func() {
				jobs <- i
				fmt.Println(<-res)
				fl := <-flag
				if fl == 1 {
					return
				}
			}()
		}
	}()
	time.Sleep(time.Second * 5)
	flag <- 1
}

func ByCtx() {
	cont, cancel := context.WithTimeout(context.Background(), time.Second*3)
	num := 20
	jnum := 100
	jobs := make(chan int, jnum)
	res := make(chan int, jnum)
	defer close(res)
	defer close(jobs)
	defer cancel()
	for i := 1; i < num; i++ {
		worker := NewWorker(i, jobs, res)
		go worker.Work()
	}
	for i := 0; ; i += 1 {
		select {
		case <-cont.Done():
			fmt.Println("DONE")
			return
		default:
			jobs <- i
			fmt.Println(<-res)
		}
	}
}

func main() {
	//ByChannels()
	ByCtx()
}
