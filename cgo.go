package main

/*
#include <stdlib.h>
#include <math.h>
typedef struct Foo {
	int a;
	int type;
}Foo;

typedef int (*intFunc) ();

int bridge_int_func(intFunc f) {
	return f();
}

int fortytwo()
{
	return 42;
}

typedef struct POINT
{
	double x;
	double y;
}POINT;
*/
import "C"
import (
	"fmt"
	"reflect"
	"time"
)

func Random() int {
	return int(C.random())
}

func Seed(i int64) {
	C.srandom(C.uint(i))
}

func Fortytwo() int {
	return 42
}

func main() {
	Seed(time.Now().Unix())
	fmt.Println(Random())

	//var foo = C.struct_Foo{65, 28}
	var foo = C.Foo{65, 28}
	//fmt.Println(foo, C.sizeof_struct_Foo)
	fmt.Println(foo, C.sizeof_Foo)
	fmt.Println(foo._type, foo.a)

	n, err := C.sqrt(-1)
	fmt.Println(n, err)

	f := C.intFunc(C.fortytwo)
	fmt.Println(reflect.TypeOf(f), int(C.bridge_int_func(f)))
	//fmt.Println(int(C.bridge_int_func(C.fortytwo))) // error

	var p C.POINT
	p.x = 9.45
	p.y = 23.12
	fmt.Println(p) // {9.45 23.12}
}
