package types

type LinuxPc struct {
	scanner base.Scanner
}

func (pc LinuxPc) Scfn() {
	println("Scan pc Linux system")
	pc.scanner.ScanFile()
}