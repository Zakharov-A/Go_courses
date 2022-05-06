package L_36

import "fmt"

type AsusMonitor struct {
	Size int
}

func (pc AsusMonitor) PrintDatails() {
	fmt.Printf("[Asus] Monitor Size:[%d]\n",pc.Size)
}

