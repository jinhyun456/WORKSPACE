package main

import "fmt"

func main() {
	var c chan int
	c = make(chan int)

	c <- 10
	v := <-c

	fmt.Println(v)
}
