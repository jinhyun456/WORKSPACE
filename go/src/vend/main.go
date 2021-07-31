// 1. 프로그램을 실행했을때 소유하고 있는 금액을 입력함
// 2. 자판기 메뉴가 출력함.
// 3. 메뉴를 번호로 입력받아 소지하고 있는 금액을 감소 시킴
// 4. 메뉴가 나왔다는 문구 출력 후 무한 반복
// 5. 소지금액이 다 소진됐을 때 무한 반복문을 탈출함

package main

import "fmt"

func main() {
	// 소지금액을 입력하기 위해선, 소지금액을 받는 변수필요
	var cost int

	_, error := fmt.Scanf("%d\n", &cost)

	if error != nil {
		fmt.Println("잘못 입력했습니다.")
	} else {
		if (cost < 3000) || (cost > 4000) {
			continue
		} else {
			break
		}
	}
}
