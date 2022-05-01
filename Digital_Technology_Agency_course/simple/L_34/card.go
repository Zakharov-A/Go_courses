package l_34

import "time"

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	println("[Card] Request to the bank to check the balance")
	time.Sleep(time.Microsecond * 800)
	return card.Bank.CheckBalance(card.Name)

}
