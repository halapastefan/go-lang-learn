package main

import (
	"fmt"
	"time"
)

const (
	Microsecond time.Duration = 1
	Nano						= 1000 * Microsecond
)

const (
	A1 = iota // 0
	B1 = iota // 1
	C1 = iota // 2
)

const (
	A2 = iota + 1 // 1
	B2 // 2
	C2 // 3
)
func main() {

	const ti = 28 //kind constant
	const tt int = 3 //type consant

	a := ti * tt

	fmt.Println(a)
}
