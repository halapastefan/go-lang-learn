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
