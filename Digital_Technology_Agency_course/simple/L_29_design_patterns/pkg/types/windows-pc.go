package types

import "simple/L_29_design_patterns/pkg/base"

type WinPc struct {
	scanner base.Scanner
}

func (pc *WinPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner

}

func (pc *WinPc) Scan() {
	println("Scan pc Windows system")
	pc.scanner.ScanFile()
}
