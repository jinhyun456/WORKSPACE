package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++ // 1초간격으로 0부터시작 1씩 증가하는 인티져(정수형)  쓰레드
	}
}

func main() {
	c := make(chan int) // channel 생성      선언을 해서 초기화 해줘야 한다  -> 채널을 대입을 시키기전에 초기화시키고 대입을 시켜라

	go push(c) //channel을 push함수에 c인자로 넘기고 go Thread시작

	timerChan := time.After(10 * time.Second)  //After 한번만 알려준다
	tickTimeChan := time.Tick(2 * time.Second) //tick 2초 간격 주기적으로 알려준다.

	for {
		select {
		case v := <-c: //c에 입력값이 있으면 아래 v출력
			fmt.Println(v)
		case <-timerChan:
			fmt.Println("timed out")
			return
		case <-tickTimeChan: //tick은 시간주기적 알림   대기하고 있다가 값이들어오면 출력
			fmt.Println("Tick")
		}
	}
}
