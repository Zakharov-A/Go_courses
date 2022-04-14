package types

import "simple/L_29_design_patterns/pkg/base"

type MacPc struct {
	scanner base.Scanner
}

func (pc *MacPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner

}

func (pc *MacPc) Scan() {
	println("Scan pc Linux system")
	pc.scanner.ScanFile()
}
