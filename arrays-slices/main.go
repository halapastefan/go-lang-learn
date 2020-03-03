package main

import "fmt"
/*
Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element
(as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents.
(To avoid the copy you could pass a pointer to the array, but then that's a pointer to an array, not an array.)
One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.
 */
func main() {

	//Possible bug

	var data = make([]string, 5)

	inspectSlice(data)
	data[0]="One"
	data[1]="two"
	data[2]="three"
	data[3]="four"
	data[4]="five"
	inspectSlice(data)

	//After appending to orignal slice, slice will get bigger, so it will reserve new block and this will point to old
	//GC cant to clean up because there is a reference to it, and that can lead to memory leak
	superData := &data[3]

	fmt.Printf("Data %s\t\t Adress: %v", *superData, superData)

	data = append(data, "new")
	inspectSlice(data)

	*superData = "new"
}

func inspectSlice(s []string) {

	fmt.Printf("Len: %d \tCap: %d \n", len(s), cap(s))
	/*
		range is using value semantics, so v in this case will be copy, and &v is different than real thing.
		When accessing via index, we access real value.
	 */
	for i, v := range s {
		fmt.Printf("Index: %d,\t value: %s,\t\t\t value addr: \t%v pointer: %v \n", i, v, &v, &s[i])
	}
}


func example03() {
	var data = make([]string, 5)

	data[0]="One"
	data[1]="two"
	data[2]="three"
	data[3]="four"
	data[4]="five"

	fmt.Println(data)

	inspectSlice(data)

	for _, v := range data {
		v = "terror"
		fmt.Printf(v)
	}
	fmt.Println("---------------------------------------------------")
	inspectSlice(data)

	fmt.Println("-------------------Slicing--------------------------")
	//slicing
	slice1 := data[1:3]
	inspectSlice(slice1)

	fmt.Println("-------------------Slicing: side effect--------------------------")
	data[1] = "CHANGED"
	inspectSlice(data)
	inspectSlice(slice1)

	fmt.Println("-------------------Slicing: no side effect--------------------------")

	slice2 := data[1:3:3]
	inspectSlice(slice2)
	slice2[0] = "CHANGED2"

	slice2 = append(slice2, "new")
	inspectSlice(slice2)
	inspectSlice(data)
}
func example02()  {
	var data = [5]string{}

	data[0] = "Stefan"
	data[1] = "Blabla"
	data[2] = "Blabla1"
	data[3] = "Blabla3"
	data[4] = "Blabla4"

	//value semantic
	for i, v := range data {

		v = "ppppp"
		fmt.Println(i, v, &v, &data[i])
	}

	for i := range data {
		fmt.Println(&data[i])
	}
	fmt.Println(data)



}
func someFunc(str string) {
	fmt.Println("In function: ", str, &str)
}

// this copies whole array
func doSomething(data [5]int) {
	fmt.Println(data)
}
// this passes array as reference
func doSomething2(data *[5]int)  {
	fmt.Println(data)
}

func example01()  {
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
