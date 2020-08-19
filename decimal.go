package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}
	fmt.Println(n)

	d := decimal.Zero
	f := decimal.NewFromFloat(.01)
	for i := 0; i < 1000; i++ {
		d = d.Add(f)
	}
	fmt.Println(d)
}
