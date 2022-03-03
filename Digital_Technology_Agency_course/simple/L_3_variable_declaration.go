package main

import "fmt"

// ---- global variable ----
// var First = "Good Word" //Глобал

// var first = "word" // Переменная с маленькой буквы внутри пакета

// ---- String output ----
// func main() {
// 	first := "Roghan the best!"
// 	Run(first)

// }

// func Run(xxx string) {
// 	fmt.Printf("%s", xxx)
// }

// ----

// ---- Hello ----
// func main() {
// 	var x string
// 	x = "Hello World"
// 	fmt.Println(x)
// }

// ----

// ---- "Addition of strings" ----
// func main() {
// 	var x string
// 	x = "first "
// 	fmt.Println(x)
// 	x = x + "second"
// 	fmt.Println(x)

// }

// ----

// ---- constant declaration ----
// func main() {
//     const x string = "Hello World"
//     fmt.Println(x)
// }

// ----

// ---- Defining Multiple Variables ----
// func main() {
// 	var (
//     a = 5
//     b = 10
//     c = 15
// )
//
// fmt.Println(a, b, c)
// }

// ----

// ---- Program example ----

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

// ----
