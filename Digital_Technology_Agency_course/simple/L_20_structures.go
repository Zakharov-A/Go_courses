package main

import "fmt"

// ---- Structure initialization ----

// type Pc struct {
// 	Core int
// 	Hard int
// }

// func main() {
// 	var p Pc = Pc{
// 		Core: 4,
// 		Hard: 250,
// 	}
// 	fmt.Printf("Core = [%d] Hard [%d]", p.Core, p.Hard)

// }

// ----

// ---- Default structure initialization ----

// type Pc struct {
// 	Core int
// 	Hard int
// }

// func main() {
// 	var p Pc = Pc{}
// 	fmt.Printf("Core = [%d] Hard [%d]", p.Core, p.Hard)

// }

// ----

// ---- Structures for changing fields ----

type Pc struct {
	Core int
	Hard int
}

func main() {
	var p Pc = Pc{}
	p.Core = 8
	p.Hard = 500
	fmt.Printf("Core = [%d] Hard [%d]", p.Core, p.Hard)

}

// ----
