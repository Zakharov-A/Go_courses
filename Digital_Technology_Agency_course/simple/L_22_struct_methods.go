package main

import "fmt"

// ---- Struct Methods ----

type Cpu struct {
	Prod string
	Core int
	Hard int
}

type Pc struct {
	Cpu
}

func (pc Pc) On() {
	println("[On] Shutdown PC")
}

func (pc Pc) Off() {
	println("[Off] Turning on PC")
}

func (pc Pc) initMethod() {
	println("[Init] ... ")
}

func main() {
	var p Pc = Pc{
		Cpu: Cpu{
			Prod: "Intel",
			Core: 8,
		},
	}
	p.On()
	p.Off()
	fmt.Printf("Pc CPU - [%v]", p.Cpu)
}

// ----
