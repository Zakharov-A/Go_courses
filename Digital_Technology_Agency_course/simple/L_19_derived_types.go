package main

import "fmt"

// ---- Derived types ----

// type km int

// func main() {
// 	var distance km = 100
// 	fmt.Println(distance)

// }

// ----

//  ---- Derived types option 2 ----

type km string

func main() {
	var distance km = "Good"
	fmt.Println(distance)

}

// ----
