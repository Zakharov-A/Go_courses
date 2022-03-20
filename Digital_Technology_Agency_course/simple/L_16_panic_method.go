package main

import "fmt"

// ---- Error output ----

func main() {
	a := 0
	b := 6
	if a == 0 {
		panic("Division by 0")
	}
	fmt.Println(b / a)
}

// ----
