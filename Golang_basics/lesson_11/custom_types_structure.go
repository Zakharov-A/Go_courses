package main

import (
	"fmt"
	"time"
)

type OurString string
type OurInt int

type Person struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	// var customString OurString
	// fmt.Printf("%T %#v \n", customString, customString)

	// customString = "Hello World"
	// fmt.Printf("%T %#v \n", customString, customString)

	// customInt := OurInt(5)
	// fmt.Printf("%T %#v \n", customInt, customInt)
	// fmt.Println(int(customInt))

	// create default
	// var John Person
	// fmt.Printf("%T, %T, %#v \n", John, John, John)

	// John = Person{}
	// fmt.Printf("%T, %T, %#v \n", John, John, John)

	// fields accessing
	// John.Name = "John"
	// John.Surname = "Wu"
	// John.Age = 38
	// fmt.Println(John)

	// create with named field
	Brad := Person{
		Name:    "Brad",
		Surname: "Simson",
		Age:     25,
	}
	fmt.Println(Brad)

	// create without field names
	Vladimir := Person{"Vladimir", "Disel", 55}
	fmt.Println(Vladimir)

	// field accessing through the pointer
	pVladimir := &Vladimir
	// fmt.Println((*pVladimir).Age)
	fmt.Println(pVladimir.Age)

	// create pointer to struct
	pIvan := &Person{"Ivan", "Drago", 45}
	fmt.Println(pIvan)

	// Anonymous structures
	unnamedStruct := struct {
		Name, LastName, BirthDate string
	}{
		Name:      "NoName",
		LastName:  "NoLastName",
		BirthDate: fmt.Sprintf("%s", time.Now()),
		// BirthDate: time.Now().String(),
	}
	fmt.Println(unnamedStruct)

}
