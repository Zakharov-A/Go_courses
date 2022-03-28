package main

import "fmt"

type Pc struct {
	Core int
	Hard int
}

func main() {
	var p Pc = Pc{}
	p.Core = 8
	p.Hard = 500
	fmt.Printf("Core = [%d] Hard [%d]", p.Core, p.Hard)

}
