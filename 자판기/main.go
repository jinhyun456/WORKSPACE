package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 돈 소지금액을 입력받는 기능
	fmt.Println("소지금액을 입력하세요.")
	// 버퍼라는 컴퓨터에 저장 메모리. (가상 메모리)
	// 가장대표적으로 복사하고 클립보드에 복사한 데이터가 들어가는데, 이 클립보드를 다른 말로는 버퍼.
	// 버퍼안에 기본적인 입출력을 담당하는 패키지가 bufio 버퍼에 데이터를 입력(담는다.) 버퍼안에 데이터를 가져오는 기능
	// BUFio .NEWREADER (버퍼 안에 데이터를 읽는 기능을 하는 기계)를 만들었다. 그 기계는 운영체제에서(윈도우 MAC) standard(기본)
	// 입출력 기능을 입력시키겠다.
	reader := bufio.NewReader(os.Stdin)
	// reader 로봇이 사용자가 키보드로 입력을 했을 때, 이 입력받은 데이터는 문자열일 것이다.
	// 그런데 이 문자열이 (엔터가쳐진) 문자열을 읽어오겠다. 두 번째에(_) 이것은 error 타입을 받는 건데,
	// _ 이렇게 빈 공간으로 받는 이유는 error 안 받겠다. 에러가나도 무시하겠다. (엔터는 (\n)
	str, _ := reader.ReadString('\n')
	// strings (문자열들)을 제어하는 패키지. 이 문자열에 존재하는 이상한 값이나 쓰레기같은 불필요없는 데이터나
	// 공백들을 모두 제거 시키겠다.
	str = strings.TrimSpace(str)
	// strconv 문자열을 다른 타입으로 변환시키는 패키지.
	// Atoi (문자열을) 정수로 변환시켜주는 함수. 입력받은 str 데이터를 정수로 만든다.
	cost, _ := strconv.Atoi(str)
	// 입력받은 문자열 그대로쓰겠다. (위에 코드가 없으면.)
	// 메뉴를 보여주는 기능
	// scanf() 포인터를 깊숙히 들어가야되요 이해하려면. (주소를 가리키는 기능) 이걸 깊숙히.
	// bufio도 완전히 이해하려면 주소에 대한 개념을 깊숙히 파헤쳐야됨.
	fmt.Println("---------------------------------------------")
	fmt.Printf("\t\t1. 콜라 500원 2. 사이다 500원 3. 밀키스 600원 4. 환타 800원 5. 솔의눈 1000원 6번 외에는 700원\n")
	fmt.Println("---------------------------------------------")
	fmt.Println("음료수를 골라주세요 : ")

	// 음료수 입력하는 기능
	res := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		drink, _ := reader.ReadString('\n')
		drink = strings.TrimSpace(drink)
		result, _ := strconv.Atoi(drink)
		res = result
		break
	}

	// 음료수를 읽는 기능.
	switch res {
	case 1:
		cost = cost - 500
		fmt.Println("콜라가 나왔습니다.")
	case 2:
		cost = cost - 500
		fmt.Println("사이다가 나왔습니다.")
	case 3:
		cost = cost - 600
		fmt.Println("밀키스가 나왔습니다.")
	case 4:
		cost = cost - 800
		fmt.Println("환타가 나왔습니다.")
	case 5:
		cost = cost - 1000
		fmt.Println("솔의 눈이 나왔습니다.")
	default:
		cost = cost - 700
		fmt.Println("아무거나 나왔습니다.")
	}
	// 음료수가 나오는 결과를 보여주는 기능.
	fmt.Printf("%v번 음료수가 나왔고, 거스름돈은 %v 입니다.", res, cost)
}
