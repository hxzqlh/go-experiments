package main

import "fmt"
import "unsafe"

func test1() {
	a := []int{}
	fmt.Println(a, len(a), cap(a))

	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 2)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 3)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 4)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 5)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 6)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 7)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 8)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 9)
}

func test2() {
	a := make([]int, 8) // [0 0 0 0 0 0 0 0]
	b := a[1:4]         // [0 0 0]

	a = append(a, 1) // [0 0 0  0 0 0 0 0 0]
	a[2] = 42        // [0 0 42 0 0 0 0 0 0]

	fmt.Println(a, len(a), cap(a)) // [0 0 42 0 0 0 0 0 1], 9, 16
	fmt.Println(b, len(b), cap(b)) // [0 0 0], 3, 7
}

func SubtractOneFromLength(slice []int) []int {
	slice = slice[0 : len(slice)-1]
	return slice
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(s, len(s), cap(s))
	s = append(s, 3)
	fmt.Println(s, len(s), cap(s))

	fmt.Println(unsafe.Sizeof(int(0)))

	a := make([]int, 49)
	fmt.Println(len(a), cap(a))

	a = append(a, 128)
	fmt.Println(len(a), cap(a))
}
