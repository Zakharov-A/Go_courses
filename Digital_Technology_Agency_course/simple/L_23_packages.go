package main

import (
	com "simple/computers"
)

// ---- Packages ----

func main() {
	var p com.Pc = com.Pc{
		Cpu: com.Cpu{
			Prod: "Intel",
			Core: 8,
		},
	}
	p.On()
	p.Off()
}

// ----
