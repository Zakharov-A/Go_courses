package main

import "fmt"

func main() {
	age := 15
	// Basic if
	if age < 18 {
		fmt.Println("You are too young (full)")
	}

	// Short syntax
	if isChild := isChildren(age); isChild == true {
		fmt.Println("You are very young (short)")
	}

	age2 := 20
	if age2 < 18 {
		fmt.Println("You are too young!")
	} else {
		fmt.Println("You are in school")
	}

}

func isChildren(age int) bool {
	return age < 18
}
