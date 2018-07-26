package main

import "fmt"

var g chan int
var quit chan chan bool

// main goroutine 让 B goroutine 先退出
func main() {
	g = make(chan int)
	quit = make(chan chan bool)
	go B()
	for i := 0; i < 3; i++ {
		g <- i
	}

	wait := make(chan bool)
	quit <- wait // 告诉 B 去退出

	<-wait // 等待B的退出
	fmt.Println("Main quit")
}

func B() {
	for {
		select {
		case i := <-g:
			fmt.Println(i + 1)
		case c := <-quit:
			fmt.Println("B quit")
			c <- true
			return
		}
	}
}
