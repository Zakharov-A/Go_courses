package l_34

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	println("[Store] request to the user to receive the balance on the card")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Shop] Checking - can [%s] buy a product! \n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Shop] Not enough funds to pay for the purchase of goods!")
		}
		fmt.Printf("[Shop] Product[%s] buy! \n", prod.Name)
	}
	return nil
}
