package main

import (
	// "errors"
	"fmt"
)

// func main() {
// 	name := "Sam"
// 	pName := &name

// 	fmt.Println(name)
// 	fmt.Println(pName)

// 	*pName = "Den"

// 	fmt.Println(name)

// }

// ---- Pointer ----

// func main() {
// 	name := "Den"
// 	fmt.Println(name)

// 	setName(&name)

// 	fmt.Println(name)

// 	var emptyPointer *string
// 	fmt.Println(emptyPointer)

// }

// func setName(name *string) {
// 	*name = "Li Sin"

// }

// ---

// ----

// const winePrice = 100

// func main() {
// 	change, err := buyWine(28, 110)
// 	if err != nil {
// 		fmt.Println("Purchase unsuccessful: ", err.Error())
// 	} else {
// 		fmt.Printf("Your change - %d, Have a nice day!\n", change)
// 	}

// 	change, err = buyWine(17, 110)
// 	if err != nil {
// 		fmt.Println("Purchase unsuccessful: ", err.Error())
// 	} else {
// 		fmt.Printf("Your change - %d, Have a nice day!", change)
// 	}

// 	change, err = buyWine(33, 98)
// 	if err != nil {
// 		fmt.Println("Purchase unsuccessful: ", err.Error())
// 	} else {
// 		fmt.Printf("Your change - %d, Have a nice day!", change)
// 	}
// }

// func buyWine(age, moneyAmount int) (int, error) {
// 	if age >= 18 && moneyAmount >= winePrice {
// 		return moneyAmount - winePrice, nil
// 	} else if age < 18 {
// 		return moneyAmount, errors.New("You are under 18 years of age, the sale of alcohol to you is prohibited!")
// 	} else {
// 		return moneyAmount, errors.New("You don't have enough money!")
// 	}
// }

// ----

// func main() {
// 	contactsList := [4]string{"Sam", "Den", "Lena", "Oliver"}

// 	fmt.Println(contactsList[0])
// 	fmt.Println(contactsList[1])
// 	fmt.Println(contactsList)

// 	fmt.Println("Array length:", len(contactsList))
// }

//  ----

// func main() {
// 	contactsList := [4]string{"Sam", "Den", "Lena", "Oliver"}
// 	for index := range contactsList {
// 		fmt.Printf("%d. %s\n", index+1, contactsList[index])
// 	}
// }

// ---- Array ----

// func main() {
// 	contactsList := [4]string{"Sam", "Den", "Lena", "Oliver"}

// 	fmt.Println("Contact list:")
// 	for index, value := range contactsList {
// 		fmt.Printf("%d. %s\n", index+1, value)
// 	}

// 	// for i := 0; i <= 10; i++ {
// 	// 	fmt.Printf("%d ", i)
// 	// }

// 	for i := 0; i < len(contactsList); i++ {
// 		fmt.Printf("%s ", contactsList[i])
// 	}
// }

// ---- Slice ----

// func main() {
// 	contactsList := []string{"Sam", "Den", "Lena", "Oliver"}

// 	fmt.Println("Contact list:")
// 	for index, value := range contactsList {
// 		fmt.Printf("%d. %s\n", index+1, value)
// 	}

// 	contactsList = append(contactsList, "Ron", "Sarah", "Lulu\n")

// 	for i := 0; i < len(contactsList); i++ {
// 		fmt.Printf("%s ", contactsList[i])
// 	}
// }

// ---- Map ----

// func main() {
// 	contactsList := map[string]string{
// 		"Sam":    "495-17-17",
// 		"Den":    "495-18-18",
// 		"Lena":   "499-19-19",
// 		"Oliver": "455-66-66",
// 	}

// 	fmt.Println("Before delete:\n ")
// 	for name, phoneNumber := range contactsList {
// 		fmt.Println(name, phoneNumber)
// 	}

// 	delete(contactsList, "Lena")

// 	fmt.Println("After delete:\n ")
// 	for name, phoneNumber := range contactsList {
// 		fmt.Println(name, phoneNumber)
// 	}

// }

// ---- Structure ----

// type contact struct {
// 	firstName   string
// 	lastName    string
// 	phoneNumber string
// 	email       string
// 	address     string
// 	dateDfBirth string
// }

// func main() {
// 	c1 := contact{
// 		firstName:   "Sam",
// 		lastName:    "Li",
// 		phoneNumber: "499-66-66",
// 		email:       "Best@best.org",
// 		address:     "London",
// 		dateDfBirth: "01/01/1900",
// 	}

// 	fmt.Println(c1)

// }

// ---- Structure ----

type contact struct {
	firstName   string
	lastName    string
	phoneNumber string
	email       string
	address     string
	dateDfBirth string
}

func (c contact) printInfo() {
	fmt.Printf("Name: %s\nSurname: %s\nPhone number: %s\nEmail: %s\nAddress: %sDate of Birth: %s\n",
		c.firstName, c.lastName, c.phoneNumber, c.email, c.address, c.dateDfBirth)
}

func (c *contact) setName(name string) {
	c.firstName = name
}

func main() {
	c1 := contact{
		firstName:   "Sam",
		lastName:    "Li",
		phoneNumber: "499-66-66",
		email:       "Best@best.org",
		address:     "London",
		dateDfBirth: "01/01/1900",
	}

	c1.setName("Rangar")

	c1.printInfo()

}
