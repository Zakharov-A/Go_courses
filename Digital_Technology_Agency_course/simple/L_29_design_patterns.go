package main

import "simple/L_29_design_patterns/pkg/types"

var (
	// scannners
	hpScanner = types.Hp{}
	epScanner = types.Epson{Name: "Sony"}
	// computers
	LinuxPc = types.LinuxPc{}
	WinPc   = types.WinPc{}
	MacPc   = types.MacPc{}
)

func main() {

	LinuxPc.AddScanner(hpScanner)
	LinuxPc.Scan()
	WinPc.AddScanner(epScanner)
	WinPc.Scan()

}
