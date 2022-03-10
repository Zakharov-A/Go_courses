package main

import "fmt"

//---- output of numbers ----

// func main() {

// 	fmt.Println("Go the Best!")
// 	fmt.Println("1")
// 	fmt.Println("2")
// 	fmt.Println("3")
// 	fmt.Println("4")
// 	fmt.Println("5")
// 	fmt.Println("6")
// 	fmt.Println("7")
// 	fmt.Println("8")
// 	fmt.Println("9")
// 	fmt.Println("10")
// }

// ----

//Функция for 1 вариант

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}
// }

// ----

//Функция for 2 вариант

// func main() {
// 	for i := 1; i <= 8; i++ {
// 		fmt.Println(i)
// 	}
// }

// ----

// ---- for function together with if and else ----

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i % 2 == 0 {
// 			fmt.Println(i, "even")
// 		} else {
// 			fmt.Println(i, "odd")
// 		}
// 	}
// }

// ----

// ---- Displaying a number to choose from a list ----

// func main() {
// 	i := 4
// 	if i == 0 {
// 		fmt.Println("Zero")
// 	} else if i == 1 {
// 		fmt.Println("One")
// 	} else if i == 2 {
// 		fmt.Println("Two")
// 	} else if i == 3 {
// 		fmt.Println("Three")
// 	} else if i == 4 {
// 		fmt.Println("Four")
// 	} else if i == 5 {
// 		fmt.Println("Five")
// 	}
// }

// ----

// Displaying a number to choose from a list
// using if switch and case

// func main() {
// 	i := 5
// 	switch i {
// 	case 0:
// 		fmt.Println("Zero")
// 	case 1:
// 		fmt.Println("One")
// 	case 2:
// 		fmt.Println("Two")
// 	case 3:
// 		fmt.Println("Three")
// 	case 4:
// 		fmt.Println("Four")
// 	case 5:
// 		fmt.Println("Five")
// 	default:
// 		fmt.Println("Unknown Number")
// 	}
// }

// ----

// ---- Printing odd numbers ----
// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// ----

// ---- if, else if----
// func main() {
// 	a := 8
// 	b := 11
// 	if a < b {
// 		fmt.Printf("a < b %v", a < b)
// 	} else if a > b {
// 		fmt.Printf("a > b")
// 	} else {
// 		fmt.Printf("Another way!")
// 	}

// }

// ----

// ---- if, else if----
func main() {
	a := 9
	switch a {
	case 5:
		fmt.Println("a = 5")
	case 8:
		fmt.Println("a = 8")
	default:
		fmt.Println("Unknown Number")
	}

}

// ----
