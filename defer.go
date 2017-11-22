package main

import (
	"fmt"
	"runtime/debug"
)

func foo() {
	defer func() {
		fmt.Println("in bar")
	}()

	panic("ops")
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			debug.PrintStack()
		}
	}()

	foo()
}
