package main

import (
	"fmt"
	"runtime"
)

type Person struct {
	Age int
	Name string
}

func (p *Person)setAge(age int)  {
	p.Age = age
}

func main() {

	p := Person{
		Age:  2,
		Name: "Stefan",
	}

	fmt.Println(p.Age)
	p.setAge(3)
	fmt.Println(p.Age)

	(&p).setAge(3)

	fmt.Println(runtime.NumCPU())
}

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}
