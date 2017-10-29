package main

import "fmt"

func test1() {
	data := make(map[string]int)
	data["one"] = 1
	data["two"] = 2
	data["three"] = 3

	for k, v := range data {
		fmt.Println(k, v)
		if v == 1 {
			data["four"] = 4
			delete(data, "three")
		}
	}

	fmt.Println(data)
}

func main() {
	test1()
}
