package main

import "fmt"

type contact struct {
  firstName string
  lastName string
  phoneNumber string
  email string
  address string
  dateDfBirth string
}

func (c contact) printInfo() {
    fmt.Printf("Имя: %s\nФамилия: %s\nНомер телефона: %s\nE-mail: %s\nАдрес: %s\nДата рождения: %s\n",
    c.firstName, c.lastName, c.phoneNumber, c.email, c.address, c.dateDfBirth)
}


func main() {
  c1 := contact{
    firstName: "Tom",
    lastName: "Huck",
    phoneNumber: "8-999-666-66-66",
    email: "nagibhabib@gmail.com",
    address: "ул. Героя Шаляпина 58а, г. Новопопкинс",
    dateDfBirth: "20.12.1980",
  }

  c1.printInfo()

}
