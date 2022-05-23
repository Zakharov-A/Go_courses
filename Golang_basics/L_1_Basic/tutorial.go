package main

import "fmt"

// ----
// func main() {
// 	var mess string
// 	mess = "Hello World"

// 	var name string
// 	var age int
// 	var weigt float32
// 	var isAdult bool

// 	weigt = 69.3
// 	age = 28

// 	fmt.Println(name)
// 	fmt.Println(age)
// 	fmt.Println(weigt)
// 	fmt.Println(isAdult)
// 	fmt.Println(mess)

// }
// ----

func main() {
	var name string
	var age int
	var weight float32

	name = "Sam"
	age = 105
	weight = 95

	printPerson(name, age, weight)
	printPerson("Den", 35, 78.7)

	fmt.Print(isAdult(15))

}

func printPerson(name string, age int, weight float32) {
	fmt.Printf("Name: %s\nAge: %d\nweigt: %.1f\n", name, age, weight)
}

func isAdult(age int) bool {
	return age >= 18

}
