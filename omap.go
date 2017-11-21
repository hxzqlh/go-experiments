package main

import "fmt"

type MapList struct {
	data map[string]interface{}
	keys []string
}

func NewMapList() *MapList {
	return &MapList{
		data: make(map[string]interface{}),
		keys: []string{},
	}
}

func (mapList *MapList) HasKey(k string) bool {
	_, exists := mapList.data[k]
	return exists
}

func (mapList *MapList) Set(key string, val interface{}) {
	if !mapList.HasKey(key) {
		mapList.keys = append(mapList.keys, key)
	}

	mapList.data[key] = val
}

func (mapList *MapList) Delete(key string) {
	if !mapList.HasKey(key) {
		return
	}

	pos := indexOf(mapList.keys, key)
	mapList.keys = remove(mapList.keys, pos)
	delete(mapList.data, key)
}

func (mapList *MapList) Size() int {
	return len(mapList.data)
}

func (mapList *MapList) Walk(cb func(key string, data interface{})) {
	for _, k := range mapList.keys {
		cb(k, mapList.data[k])
	}
}

func indexOf(data []string, word string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func remove(a []string, i int) []string {
	a[i] = a[len(a)-1]
	ret := a[:len(a)-1]
	return ret
}

func main() {
	ml := NewMapList()
	a := "ElementsAlice"
	b := 12000000000000000
	c := 3.56
	d := struct {
		Name string
		Age  int
	}{
		Name: "hxz",
		Age:  30,
	}

	ml.Set("one", a)
	ml.Set("two", b)
	ml.Set("three", c)
	ml.Set("four", d)

	cb := func(key string, data interface{}) {
		fmt.Println(key, data)
	}
	ml.Walk(cb)

	ml.Delete("one")
	ml.Walk(cb)
	fmt.Println(ml)
}
