package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2017-10-26T17:15:46.711+08:00"
	//timeStr := "2017-10-26T10:19:53.028Z"

	txTime, _ := time.Parse(time.RFC3339, timeStr)
	fmt.Println(txTime, txTime.Unix(), txTime.Location())

	loc, _ := time.LoadLocation("Asia/Chongqing")
	//txTime2, _ := time.ParseInLocation(time.RFC3339, timeStr, loc)
	//fmt.Println(txTime2, txTime2.Unix(), txTime2.Location())

	t3 := txTime.In(loc)
	fmt.Println(t3, t3.Unix(), t3.Location())
}
