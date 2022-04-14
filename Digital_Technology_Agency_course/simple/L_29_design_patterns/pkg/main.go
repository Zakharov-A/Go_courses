package main

import "simple/L_29_design_patterns/pkg/types"

var (
	// scannners
	hpScanner = types.Hp{}
	epScanner = types.Epson{}
	// computers
	linuxPc = types.LinuxPc{}
	winPc   = types.WinPc{}
	macPc   = types.MacPc{}
)

func main() {

	linuxPc.AddScanner(hpScanner)
	linuxPc.Scan()

}
