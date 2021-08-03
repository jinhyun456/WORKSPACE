package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct { //은행계좌
	balance int //잔액 (메모리 보호대상) Mutex
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) { //인출
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) { //입금
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

var accounts []*Account //slice 배열 (메모리 보호대상) Mutex

func Transfer(sender, receiver int, money int) { //송금자, 받는자, 송금액
	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)
}

func GetTotalBalance() int { //전체 잔액량
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func RandomTransfer() { //랜덤함수 레퍼런스
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 { //샌더 잔액여부 확인
			break //샌더 잔액이 있으면 멈춘다
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	PrintTotalBalance()

	for i := 0; i < 10; i++ {
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}

//음오아예
