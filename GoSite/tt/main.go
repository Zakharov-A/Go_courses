package main

import (
	"fmt"
	"go_site/template"
	"html/template"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money "+
		"equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	// bob.setNewName("Loki")
	// fmt.Fprintf(w, bob.getAllInfo())
	// fmt.Fprintf(w, `<h1>Main Text</h1>
	// <b>Hello!!!</b><br>
	// <b>Are you here?!</b>`)
	tmpl, _ := template.ParseFiles("template/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All contacts here!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)

}

func main() {
	// var bob User = ....
	// bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}

	handleRequest()

}
