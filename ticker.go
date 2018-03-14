package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(3 * time.Second)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case a, ok := <-ticker.C:
				fmt.Println("ticker:", a, ok)
			case <-quit:
				fmt.Println("exit ticker")
				ticker.Stop()
				wg.Done()
				return
			default:
				// Do other stuff
			}
		}
	}()
	wg.Add(1)

	go func() {
		for {
			select {
			case a, ok := <-timer.C:
				fmt.Println("time out:", a, ok)
				quit <- true
				wg.Done()
				return
			default:
			}
		}
	}()
	wg.Add(1)

	wg.Wait()
}
