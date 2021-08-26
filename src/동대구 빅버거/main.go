package main

import "fmt"

type TongsOfPatty interface {
	String() string
}

type Patty interface {
	GetOneTongs() TongsOfPatty
}

type BigBurger struct {
	val string
}

func (b *BigBurger) PutPatty(patty Patty) {
	tongs := patty.GetOneTongs()
	b.val += tongs.String()
}

func (b *BigBurger) String() string {
	return "동대구" + b.val
}

type ChickenPatty struct {
}

func (j *ChickenPatty) GetOneTongs() TongsOfPatty {
	return &TongsOfChickenPatty{}
}

type SlicePatty struct {
}

func (j *SlicePatty) GetOneTongs() TongsOfPatty {
	return &TongsOfSlicePatty{}
}

type CheesePatty struct {
}

func (j *CheesePatty) GetOneTongs() TongsOfPatty {
	return &TongsOfCheesePatty{}
}

type TongsOfChickenPatty struct {
}

func (s *TongsOfChickenPatty) String() string {
	return "+ chickenPatty"
}

type TongsOfSlicePatty struct {
}

func (s *TongsOfSlicePatty) String() string {
	return "+ SlicePatty"
}

type TongsOfCheesePatty struct {
}

func (s *TongsOfCheesePatty) String() string {
	return "+ CheesePatty"
}

func main() {
	BigBurger := &BigBurger{}
	patty := &ChickenPatty{}
	BigBurger.PutPatty(patty)

	fmt.Println(BigBurger)
}

//
