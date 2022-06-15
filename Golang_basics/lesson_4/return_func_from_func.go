package main

import "fmt"

func main() {

	divideBy2 := createDivider(2)
	divideBy10 := createDivider(10)

	fmt.Println(divideBy2(100))
	fmt.Println(divideBy10(100))

}

func createDivider(divider int /* 2 */) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider /* 2 */
	}
	return dividerFunc
}
