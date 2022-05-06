package main

import "simple/L_36"

var (
	brands = []string{L_36.Asus, L_36.Hp, "dell"}
)

func main() {
	for _, brand := range brands{
	factory, err := L_36.GetFactory(brand)
		if err != nil {
			println(err.Error())
			continue
		}
		monitor := factory.GetMonitor()
		monitor.PrintDatails()
		computer := factory.GetComputer()
		computer.PrintDatails()
	}
}

