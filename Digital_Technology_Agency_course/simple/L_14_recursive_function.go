package main

import "fmt"

// factorial 1*2*3*4

func main() {
	fmt.Print(factorial(4))

}

func factorial(f uint) uint {
	if f == 0 {
		return 1
	}
	return f * factorial(f-1)
}
