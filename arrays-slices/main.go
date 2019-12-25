package main

import "fmt"
/*
Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element
(as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents.
(To avoid the copy you could pass a pointer to the array, but then that's a pointer to an array, not an array.)
One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.
 */
func main() {

	var a = [5]int{}

	doSomething(a)
	doSomething2(&a)

	b := [2]string{"one", "two"}
	fmt.Println(b)

	c := [...]string{"one", "two", "three"}
	fmt.Println(c)

	//slice
	var s []byte
	s = make([]byte, 5, 5)
	fmt.Println(cap(s))

	ss := []string{"one", "two", "three", "four"}
	p := ss[1:3]
	fmt.Println(p)
	p[0] = "sdkjlkj"
	fmt.Println(ss)
	fmt.Println(p)
}

// this copies whole array
func doSomething(data [5]int) {
	fmt.Println(data)
}
// this passes array as reference
func doSomething2(data *[5]int)  {
	fmt.Println(data)
}
