package main

import "fmt"

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

func main() {
	num := 20
	jnum := 100
	jobs := make(chan int, jnum)
	res := make(chan int, jnum)
	for i := 1; i < num; i++ {
		worker := NewWorker(i, jobs, res)
		go worker.Work()
	}
	for i := 0; i <= jnum; i += 1 {
		jobs <- i
	}
	close(jobs)
	for i := 0; i <= jnum; i++ {
		fmt.Println(<-res)
	}
	close(res)
}
