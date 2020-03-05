package main

import "fmt"

type Data struct {
	name string
	age int
}

type bill Data

func (d Data) dispalyName()  {
	fmt.Println("My name is: ", d.name)
}

func (d *Data) setAge(age int)  {
	d.age = age
}