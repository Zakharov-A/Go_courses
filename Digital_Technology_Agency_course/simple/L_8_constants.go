package main

import "fmt"

// ---- Constans ----

var (
	str = "Tom is smart"
)

const (
	NotChange = "\nI'm super man\n"
	BeingCool = "\nBeing smart is cool\n\n"
)

func main() {
	str = "\nTom is very smart\n"
	fmt.Printf("%s", str)
	fmt.Printf("%s", NotChange)
	fmt.Printf("%s", BeingCool)
}

// ----
