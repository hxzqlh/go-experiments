package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "# Time detect faces: 3.98987\nThe number of faces found in the first image: 1\nThe number of faces found in the second image: 1\nThe distance of two faces: 0.375449\nThe two pictures have the same face\n"
	list := strings.Split(ss, "\n")
	for _, one := range list {
		fmt.Println(one)
	}
	fmt.Println(list[len(list)-2])
}
