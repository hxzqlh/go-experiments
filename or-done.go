package main

import (
	"fmt"
	"sync"
	"time"
)

func orDone(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	for _, c := range channels {
		var once sync.Once
		go func(c <-chan interface{}) {
			select {
			case <-done:
				once.Do(func() { close(out) })
			case <-c:
				once.Do(func() { close(out) })
			case <-out:
			}
		}(c)
	}
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} { //1
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	done := make(chan interface{})
	start := time.Now() //2
	go func() {
		<-orDone(done, sig(5*time.Second), sig(1*time.Hour), sig(1*time.Minute))
		fmt.Printf("done after %v", time.Since(start))
	}()

	time.Sleep(time.Second)
	done <- true
	select {}
}
