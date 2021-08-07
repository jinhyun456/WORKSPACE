package main

import "fmt"

func main() {
	num := 5

	for i := 0; i < num; i++ {
		for j := num - 1; j > i; j-- {
			fmt.Printf(" ")
		}

		for k := 0; k < (2*i)+1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
