package main

import "fmt"

// type Cpu struct {
// 	Prod string
// 	Core int
// 	Hard int
// }

// type Pc struct {
// 	Cpu Cpu
// }

// func main() {
// 	var r = Pc{
// 		Cpu: Cpu{
// 			Prod: "Intel",
// 			Core: 16,
// 			Hard: 320,
// 		},
// 	}
// 	fmt.Printf("Pc CPU(Core) = [%d]", r.Cpu.Core)
// }

// ----

type Cpu struct {
	Prod string
	Core int
	Hard int
}

type Pc struct {
	Cpu
}

func main() {
	var r Pc = Pc{
		Cpu: Cpu{
			Prod: "Intel",
			Core: 16,
			Hard: 320,
		},
	}
	fmt.Printf("Pc: \nCPU(Core) = [%d]\nCPU(Prod) = [%s]\nCPU(Hard) = [%d]\n", r.Core, r.Prod, r.Hard)
}
