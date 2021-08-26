package main

import "fmt"

func main() {
	//총 5줄짜리 피라미드
	num := 5

	for i := 0; i < num; i++ {

		//한줄씩 넘어 갈수록 앞부분 공백은 4,3,2,1,0 칸으로 줄어듬
		for j := num - 1; j > i; j-- {
			fmt.Printf(" ")
		}

		//별의 갯수는 1,3,5,7,9로 늘어남.
		//2n*1의 규칙을 가짐
		for k := 0; k < (2*i)+1; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
