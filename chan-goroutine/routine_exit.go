package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(strings <-chan string, done <-chan interface{}) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer func() {
				fmt.Println("doWork exited.")
				close(completed)
			}()

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return completed
	}

	done := make(chan interface{})
	completed := doWork(nil, done)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	v, ok := <-completed
	fmt.Println("Done.", v, ok)
}
