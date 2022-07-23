package pkg

import "time"

type Card struct {
	Name string
	Balance float64
	Bank *Bank
}

func (card Card) CheckBalance() error {
	println("[Card] Request to the bank and get balance")
	time.Sleep(time.Millisecond * 800)
	return card.Bank.CheckBalance(card.Name)
}