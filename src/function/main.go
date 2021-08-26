package main

import "fmt"

func printHello() {
	fmt.Println("Hello world!")
}

func Parameter(x int) {
	fmt.Println(x)
}

func addParameter(x int) int {
	return x + 1
}

func main() {
	printHello()

	fmt.Println("Hello World!")
	Parameter(5)
	fmt.Println(addParameter(8))
}
