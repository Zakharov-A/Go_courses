package L_36

import "fmt"

type AsusComputer struct {
	Memory int
	Cpu int
}

func (pc AsusComputer) PrintDatails() {
	fmt.Printf("[Asus] Pc Cpu:[%d] Mem:[%d]\n",pc.Cpu, pc.Memory)
}