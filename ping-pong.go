package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)

	go palyer("ping", table)
	go palyer("pong", table)

	table <- new(Ball)
	time.Sleep(time.Second)
	<-table
}

func palyer(name string, ch chan *Ball) {
	for {
		b := <-ch
		b.hits++
		fmt.Println(name, b.hits)
		time.Sleep(100 * time.Millisecond)
		ch <- b
	}
}
