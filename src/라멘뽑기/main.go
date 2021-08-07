package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var cost int
	cost = 0
	fmt.Println("물건을 집어주세요: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	switch line {
	case "연필":
		fmt.Println("연필을 집었습니다.")
		cost = 500
		fmt.Printf("%d원\n", cost)
	case "공책":
		fmt.Println("공책을 집었습니다.")
		cost = 2000
		fmt.Printf("%d원\n", cost)
	case "빗자루":
		fmt.Println("빗자루를 집었습니다.")
		cost = 3200
		fmt.Printf("%d원\n", cost)
	case "쓰레기통":
		fmt.Println("쓰레기통을 집었습니다.")
		cost = 4000
		fmt.Printf("%d원\n", cost)
	case "라면":
		fmt.Println("라면을 집었습니다.")
		cost = 800
		fmt.Printf("%d원\n", cost)
	default:
		fmt.Printf("%s를 집었습니다.\n", line)
	}
	rand.Seed(time.Now().UnixNano())
	if cost == 0 {
		first := rand.Intn(9) + 1
		second := rand.Intn(9)
		str := strconv.Itoa(first) + strconv.Itoa(second) + "00"
		cost, _ = strconv.Atoi(str)
		cost = 1000
		fmt.Printf("%d원\n", cost)
	}

	fmt.Println("얼마를 내야할까요? :")

	line01, _ := reader.ReadString('\n')
	line01 = strings.TrimSpace(line01)
	money, _ := strconv.Atoi(line01)

	if money > cost {
		money -= cost
		fmt.Printf("거스름돈 %d원입니다.", money)
	} else {
		fmt.Println("잔액이 부족합니다.")
	}
}
