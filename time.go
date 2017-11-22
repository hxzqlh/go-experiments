package main

import (
	"fmt"
	"time"

	cmn "dev.33.cn/33/common"
)

func test1() {
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

func main() {
	t := cmn.ToCstTime("2006-01-02 15:04:05", "2017-11-03 14:22:12")
	fmt.Println(t, t.Unix())
}
