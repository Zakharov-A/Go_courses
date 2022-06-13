package main

import "fmt"

func main() {

	fmt.Println()

	// first, second := 3, 6

	// var multiplier func(x, y int) int
	// multiplier = func(x, y int) int { return x * y }
	// fmt.Println(multiplier(first, second))

	// divider := func(x, y int) int { return x / y }
	// fmt.Println(divider(second, first))

	// sumFunc := func(x, y int) int {
	// 	return x + y
	// }

	// subtractFunc := func(x, y int) int {
	// 	return x - y
	// }

	// fmt.Println(calculate(first, second, sumFunc))
	// fmt.Println(calculate(second, first, subtractFunc))

	divideBy2 := createDivider(2)
	divideBy10 := createDivider(10)

	fmt.Println(divideBy2(100))
	fmt.Println(divideBy10(100))

}

// func calculate(x, y int, action func(x, y int) int) int {
// 	return action(x, y)
// }

func createDivider(divider int /* 2 */) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider /* 2 */
	}
	return dividerFunc
}
