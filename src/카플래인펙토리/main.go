package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string //Plane 만들기
}

func MakeTire(carChan chan Car, PlaneChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) { // channel 4개
	for { //무한루프로 만들기
		select { //타이어 공장
		case car := <-carChan:
			car.val += "Tire_C, " //카형 타이어
			outCarChan <- car
		case Plane := <-PlaneChan:
			Plane.val += "Tire_P, " //비행기형 타이어
			outPlaneChan <- Plane   //출력함수
		}
	}
}

func MakeEngine(carChan chan Car, PlaneChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for { //무한루프 만들기
		select {
		case car := <-carChan:
			car.val += "Engine_C, " //자동차형 엔진
			outCarChan <- car
		case Plane := <-PlaneChan:
			Plane.val += "Engine_P, " //비행기형 엔진
			outPlaneChan <- Plane
		}
	}
}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)} //int를 string으로 변환함수
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i)} //int를 string으로 변환함수
		i++
	}
}

func main() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	PlaneChan1 := make(chan Plane)
	PlaneChan2 := make(chan Plane)
	PlaneChan3 := make(chan Plane)

	go StartCarWork(carChan1) //StartCarWork은 destination 목적지로 가라!  인수carChan1 은 주소값
	go StartPlaneWork(PlaneChan1)
	go MakeTire(carChan1, PlaneChan1, carChan2, PlaneChan2)
	go MakeEngine(carChan2, PlaneChan2, carChan3, PlaneChan3)

	for {
		select {
		case result := <-carChan3:
			fmt.Println(result.val)
		case result := <-PlaneChan3:
			fmt.Println(result.val)
		}
	}
}
