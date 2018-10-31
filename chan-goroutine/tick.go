package main

import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	fmt.Println("=====")
	var x int
	var ok bool

	for i := 1; i <= 7; i++ {
		x, ok = <-requests
		fmt.Println(x, ok)
	}

	/*
		limiter := time.Tick(time.Millisecond * 200)

		for req := range requests {
			<-limiter
			fmt.Println("request", req, time.Now())
		}

		fmt.Println("ok")
	*/
}
