package main

import "fmt"

// ---- Slice initialization ----

// var integers = []int{1, 2, 3, 4}

// func main() {
// 	fmt.Printf(`%v`, integers)

// }

// ----

// ---- Slice initialization with make ----

// var integers = make([]int, 4)

// func main() {
// 	fmt.Printf(`%v`, integers)

// }

// ----

// ---- Array slice ----

// var integers = []int{1, 2, 3, 4, 5}

// func main() {
// 	small := integers[0:3]

// 	fmt.Printf(`%v`, small)
// }

// ----

// ---- Adding an element to a slice ----

var integers = []int{1, 2, 3, 4, 5}

func main() {
	small := integers[0:3]
	small = append(small, 16, 66)

	fmt.Printf(`%v`, small)
}

// ----
