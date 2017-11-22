package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	cmn "dev.33.cn/33/common"
)

type Keyer interface {
	GetKey() string
}

type MapList struct {
	sync.RWMutex
	dataMap  map[string]*list.Element
	dataList *list.List
}

func NewMapList() *MapList {
	return &MapList{
		dataMap:  make(map[string]*list.Element),
		dataList: list.New(),
	}
}

func (mapList *MapList) Exists(data Keyer) bool {
	mapList.RLock()
	defer mapList.RUnlock()

	_, exists := mapList.dataMap[data.GetKey()]
	return exists
}

func (mapList *MapList) Push(data Keyer) bool {
	if mapList.Exists(data) {
		return false
	}

	mapList.Lock()
	elem := mapList.dataList.PushBack(data)
	mapList.dataMap[data.GetKey()] = elem
	mapList.Unlock()

	return true
}

func (mapList *MapList) Remove(data Keyer) {
	if !mapList.Exists(data) {
		return
	}

	mapList.Lock()
	mapList.dataList.Remove(mapList.dataMap[data.GetKey()])
	delete(mapList.dataMap, data.GetKey())
	mapList.Unlock()
}

func (mapList *MapList) Size() int {
	mapList.RLock()
	defer mapList.RUnlock()

	return mapList.dataList.Len()
}

func (mapList *MapList) Get(key string) interface{} {
	mapList.RLock()
	defer mapList.RUnlock()

	e, ok := mapList.dataMap[key]
	if ok {
		return e.Value
	} else {
		return nil
	}
}

func (mapList *MapList) Walk(cb func(data Keyer)) {
	mapList.RLock()
	defer mapList.RUnlock()

	for elem := mapList.dataList.Front(); elem != nil; elem = elem.Next() {
		cb(elem.Value.(Keyer))
	}
}

type Balance struct {
	Currency int32
	Active   int64
	Frozen   int64
}

func (b *Balance) GetKey() string {
	return cmn.IntToString(b.Currency)
}

func read(ml *MapList) {
	cb := func(data Keyer) {
		fmt.Println(data)
	}

	for i := 0; i < 100000; i++ {
		fmt.Println(i, "size", ml.Size())
		ml.Walk(cb)
		fmt.Println(i, ml.Get("444"))
	}
}

func write(ml *MapList) {
	a := &Balance{1, 100, 200}
	b := &Balance{2, 210, 340}
	c := &Balance{3, 30, 400}
	d := &Balance{4, 344, 500}
	for i := 0; i < 100000; i++ {
		ml.Push(a)
		ml.Push(b)
		ml.Push(c)

		ml.Remove(b)
		ml.Remove(d)
	}
}

func main() {
	ml := NewMapList()

	go read(ml)
	go write(ml)

	time.Sleep(time.Second)
}
