package main

import "fmt"

func main() {
	age := 40
	// &&
	if age >= 6 && age <= 18 {
		fmt.Println("You are in school")
	}

	// ||
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("You heve to get new passport")
	}

	if !isChildren(age) {
		fmt.Println(("You are am adult"))
	}

}
func isChildren(age int) bool {
	return age < 18
}
