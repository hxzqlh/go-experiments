package main

/*
#include "foo.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.count)
	C.foo()
}
