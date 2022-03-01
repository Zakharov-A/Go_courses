package main

import "fmt"

// ---- card announcement ----

// var users = map[string]int{
// 	"Admin":   1,
// 	"Manager": 2,
// }

// func main() {
// 	fmt.Printf(`%v`, users)

// }

// ----

// ---- change by key ----

// var users = map[string]int{
// 	"Admin":   1,
// 	"Manager": 2,
// 	"Tom":     4,
// }

// func main() {
// 	fmt.Printf("Tom = %d \n", users[`Tom`])
// 	users[`Tom`] = 20
// 	fmt.Printf("Tom = %d ", users[`Tom`])

// }

// ----

// ---- for loop in map ----

// var users = map[string]int{
// 	"Admin":   1,
// 	"Manager": 2,
// 	"Tom":     4,
// 	"John":    5,
// }

// func main() {
// 	for key, value := range users {
// 		fmt.Printf("key = %s value = %d\n", key, value)
// 	}
// }

// ----

// ---- Adding an element to the map and removing it ----

var users = map[string]int{
	"Admin":   1,
	"Manager": 2,
	"Tom":     4,
	"John":    5,
}

func main() {
	users["Bob"] = 100

	for key, value := range users {
		fmt.Printf("key = %s value = %d\n", key, value)
	}
	fmt.Println("deleting a user \n")
	delete(users, "Bob")
	for key, value := range users {
		fmt.Printf("key = %s value = %d\n", key, value)
	}
}

// ----
