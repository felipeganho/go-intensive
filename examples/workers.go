package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	qtWorkers := 10000

	// inicializa os workers
	for i := 0; i < qtWorkers; i++ {
		go worker(i, ch)
	}

	// joga a carga para os workers
	for i := 0; i < 1000000; i++ {
		ch <- i
	}
}
