package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = new(sync.WaitGroup)

func foo(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finish", i, ctx.Err())
			wg.Done()
			return
		case <-time.After(time.Second):
			fmt.Println("msg from", i)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go foo(ctx, i)
		wg.Add(1)
	}

	wg.Wait()
}
