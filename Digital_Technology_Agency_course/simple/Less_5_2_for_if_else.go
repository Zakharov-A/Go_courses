package main

import "fmt"

func main() {

	// fmt.Println("Go the Best!")
	// fmt.Println("1")
	// fmt.Println("2")
	// fmt.Println("3")
	// fmt.Println("4")
	// fmt.Println("5")
	// fmt.Println("6")
	// fmt.Println("7")
	// fmt.Println("8")
	// fmt.Println("9")
	// fmt.Println("10")

	//Функция for 1 вариант

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//Функция for 2 вариант

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	//Функция for совместно с if и else if

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "even")
	// 	} else {
	// 		fmt.Println(i, "odd")
	// 	}
	// }
	// Displaying a number to choose from a list
	// using if
	// i := 3
	// if i == 0 {
	// 	fmt.Println("Zero")
	// } else if i == 1 {
	// 	fmt.Println("One")
	// } else if i == 2 {
	// 	fmt.Println("Two")
	// } else if i == 3 {
	// 	fmt.Println("Three")
	// } else if i == 4 {
	// 	fmt.Println("Four")
	// } else if i == 5 {
	// 	fmt.Println("Five")
	// }

	// Displaying a number to choose from a list
	// using if switch and case

	i := 5

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
