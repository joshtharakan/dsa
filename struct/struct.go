package main

import (
	"fmt"
)

type Foo struct {
	Name string
	Age  int
}

type Bar struct {
	Foo
}

func (f Foo) Double() int {
	return f.Age * 2
}

type Doublable interface {
	Double() int
}

func doDouble(d Doublable) {
	fmt.Println(d.Double())
}

func main() {
	f := Foo{Name: "Foo", Age: 20}
	b := Bar{Foo: f}
	fmt.Println(b)
	doDouble(f)
	doDouble(b)
}
