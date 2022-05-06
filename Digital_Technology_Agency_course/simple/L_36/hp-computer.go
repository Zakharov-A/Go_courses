package L_36

import "fmt"

type HpComputer struct {
	Memory int
	Cpu int
}

func (pc HpComputer) PrintDatails() {
	fmt.Printf("[HP] Pc Cpu:[%d] Mem:[%d]\n",pc.Cpu, pc.Memory)
}