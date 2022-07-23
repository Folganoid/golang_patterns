package main

import (
	"./pkg"
	"fmt"
)

var (
	bank = pkg.Bank {
		Name: "Monobank",
		Cards: []pkg.Card{},
	}
	card1 = pkg.Card{
		Name: "CRD-1",
		Balance: 200,
		Bank: &bank,
	}
	card2 = pkg.Card{
		Name: "CRD-2",
		Balance: 5,
		Bank: &bank,
	}
	user1 = pkg.User{
		Name: "Buyer-1",
		Card: &card1,
	}
	user2 = pkg.User{
		Name: "Buyer-2",
		Card: &card2,
	}

	prod = pkg.Product{
		Name: "Wine",
		Price: 150,
	}

	shop = pkg.Shop{
		Name: "Shop-1",
		Products: []pkg.Product{
			prod,
		},
	}

)

func main() {
	println("[Bank] cards was created")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("[%s]", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}

}