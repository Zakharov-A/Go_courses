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

	// if example
	if value == 1 {
		fmt.Println("Number is one")
	} else if value == 2 || value == 3 {
		fmt.Println("Number is two or three")
	} else if value == getFour() {
		fmt.Println("Number is four")
	} else {
		fmt.Println("Default case is shown")
	}

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

}

func getFour() int {
	fmt.Println("getFour called")
	return 4
}
