package l_34

import "time"

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	println("[Store] request to the user to receive the balance on the card")
	time.Sleep(time.Microsecond * 500)
}
