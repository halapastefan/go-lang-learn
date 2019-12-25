package main

import (
	"errors"
	"fmt"
)

func executePanic() {
	defer printMessage()
	panic("This is panic sitiation")
	fmt.Println("The function executes completely")
}

func main() {
	defer printMessage()
	executePanic()
	fmt.Println("Main block is executed completely")
}
// --------------------- Example 2 -----------------------------

func main2() {
	defer printMessage()
	fmt.Println("This is line one")
	fmt.Println("This is line Two")
	fmt.Println("this is line three")
}

func printMessage() {
	fmt.Println("This is simple defer function execution")
}
// --------------------- Example 1 -----------------------------
func main1() {
	areaValue, err := calculateArea(-1)

	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(areaValue)
}

func calculateArea(radius int) (int, error) {
	if radius < 0 {
		return 0, errors.New("provide positive number")
	}
	return radius * radius, nil
}