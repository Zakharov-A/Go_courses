package main

import (
	"fmt"
	"net/http"
)

// func index(w http.ResponseWriter, r *http.Request){
//   t, err := template.ParseFiles("templates/index.html")

//   if err != nil {
//     fmt.Fprintf(w, err.Error())
//   }
//   t.ExecuteTemplate(w, "index")
// }

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is super ease!")
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
	handleRequest()

}
