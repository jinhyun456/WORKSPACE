package main

import "fmt"

type TongsOfdish interface {
	String() string
}

type sashimi interface {
	GetOnedish() TongsOfdish
}

type joseonsashimi struct {
	val string
}

func (b *joseonsashimi) Putsashimi(patty sashimi) {
	tongs := patty.GetOnedish()
	b.val += tongs.String()
}

func (b *joseonsashimi) String() string {
	return "joseon sashimi " + b.val
}

type eelsashimi struct {
}

func (j *eelsashimi) GetOnedish() TongsOfdish {
	return &TongsOfeelsashimi{}
}

type Slicesashimi struct {
}

func (j *Slicesashimi) GetOnedish() TongsOfdish {
	return &TongsOfSlicesashimi{}
}

type salmonsashimi struct {
}

func (j *salmonsashimi) GetOnedish() TongsOfdish {
	return &TongsOfsalmonsashimi{}
}

type TongsOfeelsashimi struct {
}

func (s *TongsOfeelsashimi) String() string {
	return "+ eelsashimi"
}

type TongsOfSlicesashimi struct {
}

func (s *TongsOfSlicesashimi) String() string {
	return "+ Slicesashimi"
}

type TongsOfsalmonsashimi struct {
}

func (s *TongsOfsalmonsashimi) String() string {
	return "+ salmonsashimi"
}

func main() {
	joseonsashimi := &joseonsashimi{}
	sashimi := &salmonsashimi{}
	joseonsashimi.Putsashimi(sashimi)

	fmt.Println(joseonsashimi)

}
