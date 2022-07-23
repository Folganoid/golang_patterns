package pkg

import (
	"fmt"
	"time"
	"errors"
)

type Shop struct {
	Name string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	println("[Shop] request to the user and get balance")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Printf("[Shop] Check - can user [%s] to buy product \n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _,prod := range shop.Products {
		if product != prod.Name {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Shop] too low money")
		}
		fmt.Printf("[Shop] Product [%s] - was bought\n", prod.Name)
	}

	return nil
}