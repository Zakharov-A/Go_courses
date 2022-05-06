package L_36

import "fmt"

type HpMonitor struct {
	Size int
}

func (pc HpMonitor) PrintDatails() {
	fmt.Printf("[HP] Monitor Size:[%d]\n",pc.Size)
}

