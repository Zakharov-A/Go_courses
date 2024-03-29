package L_39

import "fmt"

type Subscriber struct {
	Name string
}

func (cons *Subscriber) Update(pubName string) {
	fmt.Printf("Sending to subscribe %s from publisher %s\n", cons.GetName(), pubName)
}

func (cons *Subscriber) GetName() string {
	return cons.Name
}
