package l_34

User struct {
	Name string
	Card *Card
}

func (user User) GetBalance() float64 {
	return user.Card.Balance
}