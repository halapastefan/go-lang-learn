package main

import "fmt"

func main() {

	colors := map[string]string{
		"R": "RED",
		"G": "GREEN",
		"B": "BLUE",
	}

	fmt.Println(colors)

	//creaate empty map

	var a map[string]string
	b := make(map[string]string)

	fmt.Println(a)
	fmt.Println(b)

	b["P"] = "W"
	fmt.Println(b)
	delete(b, "P")
	fmt.Println(b)

	//iterate over map
	printMap(colors)
}

func printMap(somtehing map[string]string) {
	for key, value := range somtehing {
		fmt.Println(key, " - ", value)
	}
}
