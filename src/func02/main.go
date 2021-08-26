package main

import "fmt"

func add(n ...int) {
	sum := 0
	for _, i := range n {
		sum += i
		fmt.Println(sum)
	}
}

func sevaraIadd(x, y int) (int, int) {
	return x + 1, y + 2
}

func emptyNamedReturn(sum int) (x, y int) {
	x = (sum * 8) + 15
	y = (sum + 20) / 2
	return
}

func main() {
	add(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(sevaraIadd(5, 8))
	fmt.Println(emptyNamedReturn(15))
}
