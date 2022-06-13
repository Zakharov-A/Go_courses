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
	first, second, third := 3, 2, 6
	first1, second1 := 23, 45
	Greet()
	PersonGreet("Sam")
	PersonGreet("Si")
	FioGreet("John", "Smith")
	FioGreetNum("John", 66)
	sum := Sum(first, second)
	fmt.Println(sum)
	summa, multiply, subtraction := SumAndMultiply(first, second, third)
	fmt.Println(summa, multiply, subtraction)
	_, multiply64 := namedSumAndMultiply(first1, second1)
	fmt.Println(multiply64)
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

func FioGreetNum(name string, surname int) {
	fmt.Printf("Hi %s %d\n", name, surname)
}

func Sum(first, second int) int {
	sum := first * second
	return sum
}

func SumAndMultiply(first, second, third int) (int, int, int) {
	return first + second, first * second, third - second
}

// func namedSumAndMultiply(first, second int) (sum1 int64, multiply1 int64) {
// 	sum1 = int64(first + second)
// 	multiply1 = int64(first) * int64(second)
// 	return
// }

func namedSumAndMultiply(first, second int) (sum1 int64, multiply1 int64) {
	sum1 = int64(first + second)
	multiply1 = int64(first) * int64(second)
	return sum1, multiply1
}
