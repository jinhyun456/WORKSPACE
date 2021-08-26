package main

import "fmt"

//type SpoonOfJam interface {
//String() string //String() string리턴타입형의 관계로 SpoonOfJam으로 표현된다
//}

type SpoonOfJam interface {
	String() string
}

type Jam interface { //오로지 관계만 선언해본다
	GetOneSpoon() SpoonOfJam //SpoonOfJam 인터페이스로 해야 종속되지 않고 반환된다
}

type Bread struct { //Bread 객체를 만든다
	val string
}

func (b *Bread) PutJam(jam Jam) { //오렌지잼이든 스토우베리 잼이든 상관없다.
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string { //Bread 메서드2 String
	return "bread " + b.val
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam { //SpoonOfJam
	return &SpoonOfStrawberryJam{}
}

//
type melonJam struct {
}

func (j *melonJam) GetOneSpoon() SpoonOfJam { //SpoonOfJam
	return &SpoonOfmelonJam{}
}

//
type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam { //SpoonOfJam
	return &SpoonOfOrangeJam{}
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

//
type SpoonOfmelonJam struct {
}

func (s *SpoonOfmelonJam) String() string {
	return "+ melon"
}

//
type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	// jam := &OrangeJam{}
	jam := &melonJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
