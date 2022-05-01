package main

import (
	"fmt"
	l_34 "simple/L_34"
)

var (
	bank = l_34.Bank{
		Name:  "Bank",
		Cards: []l_34.Card{},
	}
	card1 = l_34.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = l_34.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user = l_34.User{
		Name: "Buyer 1",
		Card: &card1,
	}
	user2 = l_34.User{
		Name: "Buyer-2",
		Card: &card2,
	}
	prod = l_34.Product{
		Name:  "Cheese",
		Price: 150,
	}
	shop = l_34.Shop{
		Name: "SHOP",
		Products: []l_34.Product{
			prod,
		},
	}
)

func main() {
	println("[Bank]  Issuance of cards.")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]\n", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
}
