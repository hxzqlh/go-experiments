package main

import "fmt"

func test1() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v) // [1 2 3 0 1 2]
}

func test2() {
	v := [3]int{1, 2, 3}
	vv := v[:]
	for i := range v {
		vv = append(vv, i)
	}
	vv[0] = 11
	fmt.Println(v, vv) // [1 2 3] [1 2 3 0 1 2]
}

func test3() {
	data := [3]int{10, 20, 30}
	fmt.Println(&data[0], &data[1], &data[2])
	for i, x := range data {
		// copy array, the 'data' in 'range data' is copy of outside 'data'
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		// modify outside 'data'
		fmt.Printf("%d %d \n", x, data[i])
	}
	fmt.Println(data)
}

func test4() {
	data := [3]int{10, 20, 30}
	fmt.Println(&data[0], &data[1], &data[2])
	for i, x := range data[:] {
		// copy slice, the 'data' in 'range data' is referenced from outside 'data'
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("%d %d %p \n", x, data[i], &data[i])
	}
	fmt.Println(data)
}

func main() {
	// test1()
	// test2()
	test3()
	test4()
}
