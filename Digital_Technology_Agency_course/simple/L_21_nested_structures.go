package main

import "fmt"

// ---- nested structures ----

// type Cpu struct {
// 	Prod string
// 	Core int
// }

// type Pc struct {
// 	Cpu Cpu
// }

// func main() {
// 	var p Pc = Pc{
// 		Cpu: Cpu{
// 			Prod: "Intel",
// 			Core: 8,
// 		},
// 	}
// 	fmt.Printf("PC CPU = [%v]", p.Cpu)

// }

// ----

// ---- displaying part of the information ----

// type Cpu struct {
// 	Prod string
// 	Core int
// }

// type Pc struct {
// 	Cpu Cpu
// }

// func main() {
// 	var p Pc = Pc{
// 		Cpu: Cpu{
// 			Prod: "Intel",
// 			Core: 8,
// 		},
// 	}
// 	fmt.Printf("PC CPU(Core) = [%v]", p.Cpu.Core)
// 	fmt.Printf("\nPC CPU(Prod) = [%s]", p.Cpu.Prod)

// }

// ----

// ---- code reduction ----

type Cpu struct {
	Prod string
	Core int
}

type Pc struct {
	Cpu
}

func main() {
	var p Pc = Pc{
		Cpu: Cpu{
			Prod: "Intel",
			Core: 8,
		},
	}
	fmt.Printf("PC CPU(Core) = [%v]", p.Core)
	fmt.Printf("\nPC CPU(Prod) = [%s]", p.Prod)

}

// ----
