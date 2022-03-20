package main

import "fmt"

// ---- Pointer creation ----

// func main() {
// 	x := 99
// 	p := &x
// 	fmt.Printf("address: %v", p)
// 	fmt.Printf("\nMeaning: %v", *p)
// 	fmt.Printf("\naddress: = %v Meaning: = %v", p, *p)
// }

// ----

// ---- Pointer change ----

// func main() {
// 	x := 99
// 	p := &x
// 	fmt.Printf("address: %v", p)
// 	fmt.Printf("\nMeaning: %v", *p)
// 	fmt.Printf("\naddress: = %v Meaning: = %v", p, *p)
// 	*p = 126
// 	fmt.Printf("\naddress: = %v Meaning: = %v", p, *p)
// }

// ----

// ---- Initializing a variable with a pointer ----

func main() {
	x := new(int)
	fmt.Printf("\naddress: = %v Meaning: = %v", x, *x)
	*x = 126
	fmt.Printf("\naddress: = %v Meaning: = %v", x, *x)
}

// ----
