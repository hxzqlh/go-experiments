package main

import (
	"fmt"
	"os"
	"time"
)

const NUM = 4

var outfile [NUM]*os.File

func main() {
	cha := make(chan struct{})
	chb := make(chan struct{})
	chc := make(chan struct{})
	chd := make(chan struct{})

	for i := 0; i < NUM; i++ {
		outfile[i], _ = os.Create(fmt.Sprintf("outfile_%d", (i + 1)))
	}

	go test(1, 0, cha, chb) // 0 ---> 3
	go test(2, 1, chb, chc) // 1 ---> 0
	go test(3, 2, chc, chd) // 2 ---> 1
	go test(4, 3, chd, cha) // 3 ---> 2

	cha <- struct{}{}
	time.Sleep(time.Second)
}

func test(i int, begIndex int, ch1, ch2 chan struct{}) {
	var nextFileIndex = begIndex
	for range ch1 {
		outfile[nextFileIndex].WriteString(fmt.Sprintf("%d ", i))
		nextFileIndex = (nextFileIndex - 1 + NUM) % NUM
		ch2 <- struct{}{}
	}
}
