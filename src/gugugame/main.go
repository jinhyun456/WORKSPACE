package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 사용자가 정답을 입력받는 변수
	var num01 int
	// 정답을 담는 변수
	var result int
	// 컴퓨터가 랜덤 숫자를 만든 숫자를 담는 변수 2개.
	var no1 int
	var no2 int
	// 점수를 담는 뱐수
	var score int

	// Seed : 씨앗, 먹이 고정값. rand패키지는 먼저 랜덤 숫자를 생성시켜주는 패키지.
	// 이 random에는 고정값 존재. 이 고정값을 정해놓으면 랜덤숫자를 정했을 때 고정값만 생성이 되요.
	// 그래서 얼핏보면 random 함수가 아닌것처럼 보이지만, 시드를 임의로 고정을 셋팅해주는 패키지가 있어요.
	// time 패키지.  time은 계속 변동.. 그러기 때문에, UnixNano라는 함수를 사용하는데,
	rand.Seed(time.Now().UnixNano())
	score = 0
	fmt.Println("구구단게임입니다. 이 게임은 구구단의 정답 맞췄을 경우 점수를 1점 얻게됩니다.")
	fmt.Println("정답을 맞히지 못했을 경우는 점수 1점이 깍이게 됩니다.")
	time.Sleep(1000)
	for {
		fmt.Println("컴퓨터가 숫자를 만들고 있습니다.")
		time.Sleep(2000)
		fmt.Println("컴퓨터는 숫자를 0~9까지 랜덤으로 숫자를 2개를 생성합니다.")
		time.Sleep(1000)
		// Intn(n) (0 <= n < 10) 0 ~ n까지 사이의 랜덤숫자를 만들어내는 함수.
		no1 = rand.Intn(10)
		no2 = rand.Intn(10)
		fmt.Println("식을 만들고 있습니다.")
		time.Sleep(1000)
		fmt.Printf("%v + %v = ? \n", no1, no2)
		result = no1 + no2
		time.Sleep(1000)
		for {
			fmt.Println("정답을 입력하세요 : ")
			_, err := fmt.Scanf("%d\n", &num01)

			if err != nil {
				fmt.Println("잘못 입력했습니다.")
				continue
			} else {
				break
			}
		}

		if num01 == result {
			fmt.Println("정답입니다.")
			score++
		} else {
			fmt.Println("틀렸습니다.")
			score--
		}

		if score == 5 {
			fmt.Println("게임에서 이기셨습니다.")
			break
		}
	}
}
