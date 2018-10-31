package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDone
}

func or2(channels ...<-chan interface{}) <-chan interface{} {
	orDone := make(chan interface{})
	for _, c := range channels {
		var once sync.Once
		go func(c <-chan interface{}) {
			select {
			case <-c:
				once.Do(func() { close(orDone) })
			case <-orDone:
				return
			}
		}(c)
	}
	return orDone
}

func or3(channels ...<-chan interface{}) <-chan interface{} {
	orDone := make(chan interface{})

	go func() {
		defer close(orDone)
		cases := []reflect.SelectCase{}
		for _, c := range channels {
			one := reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			}
			cases = append(cases, one)
		}
		reflect.Select(cases)
	}()

	return orDone
}

func or4(channels ...<-chan interface{}) <-chan interface{} { //1
	switch len(channels) {
	case 0: //2
		return nil
	case 1: //3
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() { //4
		defer close(orDone)

		switch len(channels) {
		case 2: //5
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: //6
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...): //6
			}
		}
	}()

	return orDone
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

	start := time.Now() //2
	r := <-or3(sig(2*time.Hour), sig(5*time.Minute), sig(1*time.Second), sig(1*time.Hour), sig(1*time.Minute))
	fmt.Printf("done after %v %v", time.Since(start), r) //3
}
