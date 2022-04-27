package main

import (
	"fmt"
	"simple/L_32_decorator"
)

var (
	base = L_32_decorator.BasePc{}
	home = L_32_decorator.HomePc{
		Cpu: 			4,
		GraphicsCard: 	1,
		Wrapper: 		base,
	}
	server = L_32_decorator.ServerPc{
		Cpu: 16,
		Memory: 256,
		Wrapper: base,
	}
)

func main() {
	fmt.Printf("Base Price: [%f]", base.GetPrice())
	fmt.Printf("\nHome  Cpu [%d] Cards:[%d] Price: [%f]", home.Cpu, home.GraphicsCard, home.GetPrice())
	fmt.Printf("\nServer  Cpu [%d] Memory:[%d] Price: [%f]", server.Cpu, server.Memory, server.GetPrice())
}