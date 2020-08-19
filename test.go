package main

import "fmt"

type Event struct {
	num int
}

func main() {
	ss := []*Event{}
	ss = append(ss, &Event{num: 1})
	ss = append(ss, nil)
	fmt.Println(ss)
}
