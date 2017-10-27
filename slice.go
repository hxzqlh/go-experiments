package main

import "fmt"

func main() {
	a := make([]int, 8) // [0 0 0 0 0 0 0 0]
	b := a[1:4]         // [0 0 0]

	a = append(a, 1) // [0 0 0  0 0 0 0 0 0]
	a[2] = 42        // [0 0 42 0 0 0 0 0 0]

	fmt.Println(a, len(a), cap(a)) // [0 0 42 0 0 0 0 0 1], 9, 16
	fmt.Println(b, len(b), cap(b)) // [0 0 0], 3, 7
}
