package main

import "simple/L_39"

func main() {
	sub1 := &L_39.Subscriber{Name: "Sub"}
	sub2 := &L_39.Subscriber{Name: "Sub-2"}
	sub3 := &L_39.Subscriber{Name: "Sub-3"}
	subN := &L_39.Subscriber{Name: "Sub-N"}
	chanel := L_39.Publisher{
		Name:      "Publisher channel",
		Consumers: map[string]L_39.Consumer{},
	}
	chanel.Subscribe(sub1)
	chanel.Subscribe(sub2)
	chanel.Subscribe(sub3)
	chanel.Subscribe(subN)
	chanel.Notify()
	println("Unsubscribe")
	chanel.UnSubscribe(sub3)
	chanel.Notify()
}
