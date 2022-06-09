package main

import "fmt"

/* Very nice
	Whot you do next?
Do you new it?*/
/*
func funcName(params...) (returned values) {
	function's body
}
*/

func main() {
	first, second := 2, 2
	Greet()
	PersonGreet("Sam")
	FioGreet("John", "Smith")
	sum := Sum(first, second)
	fmt.Println(sum)

}

func Greet() {
	fmt.Println("Hello guys")
}

func PersonGreet(name string) {
	fmt.Printf("Hello %s\n", name)

}

func FioGreet(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}

func Sum(first, second int) int {
	sum := first * second
	return sum
}
