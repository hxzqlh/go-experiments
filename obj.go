package main

import "fmt"

type Foo struct {
	Name string
}

func (f *Foo) GetName(data map[string]int64) {
	data["one"] = 1
}

func main() {
	f := Foo{Name: "HXZ"}
	data := map[string]int64{}
	f.GetName(data)
	fmt.Println(data)
}
