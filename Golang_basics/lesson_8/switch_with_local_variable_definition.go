package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(max-min) + min

	// base swith
	switch value {
	case 1:
		fmt.Println("Number is one")
	case 2, 3:
		fmt.Println("Number is two or three")
	case getFour():
		fmt.Println("Number is four")
	default:
		fmt.Println("Default case is shown")
	}

	// switch whith local variable definition
	switch num := rand.Intn(max-min) + min; num {
	case 1:
		fmt.Println("Number is one")
	case 2, 3:
		fmt.Println("Number is two or three")
	case getFour():
		fmt.Println("Default case is four")
		fallthrough
	case 10:
		fmt.Println("Strange things happends")
	default:
		fmt.Println("Default case is shown")
	}

	// switch without condition
	switch {
	case value > 2:
		fmt.Printf(" Value %d greater than 2 \n", value)
	case value < 2:
		fmt.Printf("Value %d greater than 2 \n", value)
	default:
		fmt.Println("Value eguals 2")
	}

}

func getFour() int {
	fmt.Println("getFour called")
	return 4
}
