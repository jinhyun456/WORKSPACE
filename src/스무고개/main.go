package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)

	count := 0
	var myanswer int
	fmt.Println("1~100중에 숫자 하나를 입력하세요.")
	for num != int(myanswer) {
		count++
		fmt.Scan(&myanswer)

		if num < int(myanswer) {
			fmt.Println(myanswer, "보다 작습니다.")
		} else if num > int(myanswer) {
			fmt.Println(myanswer, "보다 큽니다.")
		} else {
			fmt.Println("정답입니다!!! 총 시도 횟수는?", count)
		}
	}
}
