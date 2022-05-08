package main

import l_37 "simple/L_37"

func main() {
	asusCollector := l_37.GetCollector("asus")
	hpCollector := l_37.GetCollector("hp")

	factory := l_37.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	factory.SetCollector(asusCollector)
	pc := factory.CreateComputer()
	pc.Print()
}
