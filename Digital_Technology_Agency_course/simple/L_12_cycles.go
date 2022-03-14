package main

import "fmt"

// ---- example 1 ----

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}
// }

// ----

// ---- example 1 ----

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("Значение: %v\n", i)
// 	}
// }

// ----

// ---- example 3 ----

// func main() {
// 	i := 0
// 	for i < 10 {
// 		fmt.Printf("Значение: %v\n", i)
// 		i++
// 	}
// }

// ----

// ---- nested loops ----
// func main() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			fmt.Printf("Значение: i=%v; j=%v\n", i, j)
// 		}

// 	}
// }

// ----

// ---- cycles in arrays if, continue ----

// func main() {
// 	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	for _, value := range arr {
// 		if value < 8 {
// 			continue
// 		}
// 		fmt.Printf("Значение: value=%v\n", value)

// 	}
// }

// ----

// ---- cycles in arrays break  ----

func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range arr {
		if value == 5 {
			break
		}
		fmt.Printf("Значение: value=%v\n", value)

	}
}

// ----
