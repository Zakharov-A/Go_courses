package types

import "simple/L_29_design_patterns/pkg/base"

type LinuxPc struct {
	scanner base.Scanner
}

func (pc *LinuxPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner

}

func (pc *LinuxPc) Scan() {
	println("Scan pc Linux system")
	pc.scanner.ScanFile()
}
