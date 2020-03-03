package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fact(0))

	fmt.Println(Strong2(145))
}

func Strong2(n int) string {

	sum := 0

	for _, digit := range strconv.Itoa(n) {
		d, _ := strconv.Atoi(string(digit))
		sum += fact(d)
	}

	if sum == n {
		return "STRONG!!!!"
	} else {
		return "Not Strong !!"
	}
}

//Strong number is the number that the sum of the factorial of its digits is equal to number itself.
func Strong(n int) string {
	// your code here
	input := n
	end := false
	sum := 0
	for ! end {
		if n == 0 {
			end = true
		} else {
			digit := n % 10
			n = n/10
			sum = sum + fact(digit)
		}
	}

	fmt.Println("sum is:")
	fmt.Println(sum)
	if sum == input {
		return "STRONG!!!!"
	} else {
		return "Not Strong !!"
	}
}

func fact(n int) int {

	sum := 1
	for i:=1; i<=n; i++ {
		sum = sum*i
	}
	return sum
}
