package main

import "fmt"

//type SpoonOfJam interface {//외부 메서드 = 관계
// string() string // SpoonOfJam은 String()이라는 외부 공개 메서드}

type Jam interface { // 오로지 관계만 선언해본다
	GetOneSpoon() *SpoonOfStrawberryJam //SponnOfJam 으로 해야 관계로 종속되지 않는다
}

type Bread struct { // Bread 객체를 만든다
	val string
}

func (b *Bread) PutJam(jam Jam) { // 오렌지잼이든 스트로베리 잼이든 상관없다
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string { //Bread 메서드 2 String
	return "bread " + b.val
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() *SpoonOfStrawberryJam { //SpoonOfJam
	return &SpoonOfStrawberryJam{}
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange "
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

func main() {
	bread := &Bread{}
	jam := &StrawberryJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
