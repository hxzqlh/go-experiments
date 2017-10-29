package main

import "fmt"

func test1() {
	a := [3]int{1, 2, 3}
	b := a

	a[0] = 11
	fmt.Println(a, b)

	b[1] = 22
	fmt.Println(a, b)
}

func test2() {
	a := [3]int{1, 2, 3}
	b := &a
	a[0] = 11
	fmt.Println(a, b)

	b[1] = 22
	fmt.Println(a, b)
}

func main() {
	test1()
	test2()
}
