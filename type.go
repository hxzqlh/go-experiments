package main

func test1() {
	// x 是 unnamed types
	var x struct{ I int }

	// y 是 named type
	type Foo struct{ I int }
	var y Foo

	// x 和 x2 类型相同
	var x2 struct{ I int }

	// y 和 z 类型不同
	type Bar struct{ I int }
	//var z Bar

	x = y
	y = x
	x = x2
	x2 = x
	// z = y // error:cannot use y (type Foo) as type Bar in assignment
}

type User struct {
	Id   int
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
}

type Employee struct {
	Id         int
	User       // annonymous field
	Title      string
	Department string
}

func test2() {
	employee := new(Employee)
	employee.SetName("Jack")
}

func main() {
	test1()
	test2()
}
