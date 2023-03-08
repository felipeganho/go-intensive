package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	// go task("Tarefa 1")
	// go task("Tarefa 2")
	// task("Tarefa 3")

	canal := make(chan int) //canal de comunicacao entre as threads

	//T2
	go publish(canal)

	//T3
	go reader(canal)
	time.Sleep(time.Second)
	// //thread 2
	// go func() {
	// 	canal <- "Olá mundo!"
	// 	canal <- "Olá mundo 2!"
	// }()

	// //T1
	// fmt.Println(<-canal)
	// fmt.Println(<-canal)
}

func reader(canal chan int) {
	for x := range canal {
		fmt.Println(x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
