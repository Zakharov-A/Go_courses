package l_34

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Sprintf("[Bank] Receiving the balance on the card %s", cardNumber))
	time.Sleep(time.Millisecond * 300)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Bank] Insufficient funds!")

		}
	}
	println("[Bank] Balance!")
	return nil
}
