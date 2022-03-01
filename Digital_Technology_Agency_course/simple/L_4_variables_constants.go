package main

import "fmt"

// var globalVar string = "I'm super man! " // глобальная переменная

func main() {
	// var globalVar string = "I'm super man! " //Локальная переменная в рамках блока {}
	// var x string
	// var y string
	// x = "Hello World!!! "
	// fmt.Println(x)
	// y = "I'm best!"
	// fmt.Println(x)
	// x = x + y
	// fmt.Println(x)

	// youAre := "I'm super man! "
	// fmt.Println("Who are you?! ", youAre)
	// 	fmt.Println(globalVar)

	// Объявление нескольких переменных
	// var (
	// 	a       = 5
	// 	b       = 10
	// 	c       = 15
	// 	manyVar = 0
	// )

	// manyVar = a + b + c
	// fmt.Println(manyVar)
	//Константы
	// const (
	// 	ad = 5
	// 	bd = 10
	// 	c  = 15
	// )
	// sum := "summa"
	// lotConst := ad + bd + c
	// fmt.Println(sum, lotConst)

	//Функция Scanf

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)

}

// func f() {
// 	fmt.Println(globalVar)
// }
