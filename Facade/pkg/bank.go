package pkg

import (
	"fmt"
	"time"
	"errors"
)

type Bank struct {
	Name string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Printf("[Bank] get card %s balance", cardNumber))
	time.Sleep(time.Millisecond * 300)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Bank] balance is to low")
		}
	}
	println("[Bank] Money exists")
	return nil
}