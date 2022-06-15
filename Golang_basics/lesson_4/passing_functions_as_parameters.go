package main

import "fmt"

func main() {

	first, second := 1, 2

	sumFunc := func(x, y int) int {
		return x + y
	}

	subtractFunc := func(x, y int) int {
		return x - y
	}

	fmt.Println(calculate(first, second, sumFunc))
	fmt.Println(calculate(second, first, subtractFunc))

}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}
