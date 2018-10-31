package main

import (
	"fmt"
	"sync"
	"time"
)

type LimitRate struct {
	rate  int           //计数周期内最多允许的请求数
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //计数周期内累计收到的请求数
	lock  sync.Mutex
}

func (l *LimitRate) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count == l.rate-1 {
		for {
			now := time.Now()
			if now.Sub(l.begin) >= l.cycle {
				//速度允许范围内， 重置计数器
				l.Reset(now)
				return true
			} else {
				// wait
			}
		}
	} else {
		//没有达到速率限制，计数加1
		l.count++
		return true
	}
}

func (l *LimitRate) Set(r int, cycle time.Duration) {
	l.rate = r
	l.begin = time.Now()
	l.cycle = cycle
	l.count = 0
}

func (l *LimitRate) Reset(t time.Time) {
	l.begin = t
	l.count = 0
}

func main() {
	var wg sync.WaitGroup
	var lr LimitRate
	lr.Set(3, time.Second) // 1s内最多请求3次

	for i := 0; i < 10; i++ {
		wg.Add(1)

		fmt.Println("Create req", i, time.Now())
		go func(i int) {
			if lr.Allow() {
				fmt.Println("Respon req", i, time.Now())
			}
			wg.Done()
		}(i)

		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
