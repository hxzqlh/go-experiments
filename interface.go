package main

import (
	"fmt"
)

// interface
type Cat interface {
	Meow()
}

// impl1
type Tabby struct{}

func (*Tabby) Meow() {
	fmt.Println("Tabby meow")
}

func GetNilTabbyCat() Cat {
	var myTabby *Tabby = nil
	return myTabby
}

func GetTabbyCat() Cat {
	var myTabby *Tabby = &Tabby{}
	return myTabby
}

// impl2
type Gafield struct{}

func (*Gafield) Meow() {
	fmt.Println("Gafield meow")
}

func GetNilGafieldCat() Cat {
	var myGafield *Gafield = nil
	return myGafield
}

func GetGafieldCat() Cat {
	var myGafield *Gafield = &Gafield{}
	return myGafield
}

func main() {
	var (
		cat1 Cat = nil
		cat2     = GetNilTabbyCat()
		cat3     = GetTabbyCat()
		cat4     = GetNilGafieldCat()
	)

	fmt.Println(cat1 == nil, cat1 == cat2, cat2 == cat3, cat2 == cat4)
}
