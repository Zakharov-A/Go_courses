package main

import "fmt"

func main() {
	var a = 5
	var b = 2
	var c = 2
	// ---- relation operations ----
	fmt.Printf("\n[Equals] a == b:[%v]", a == b)        // -- equal false --
	fmt.Printf("\n[Equals] a == c:[%v]", b == c)        // -- equal true--
	fmt.Printf("\n[More] a > b:[%v]", a > b)            // -- more --
	fmt.Printf("\n[Not equal] a != b:[%v]", a != b)     // -- not equal --
	fmt.Printf("\n[More or equal] a >= b:[%v]", a >= b) // -- more or equal --
	// ----

	// ---- Boolean operations ----
	var aTrue = false
	fmt.Printf("\n[Negation] !aTrue(true) [%v]\n", !aTrue) // -- negation --
	var g bool = 5 > 4
	var h bool = 6 > 8

	fmt.Printf("\n[Conjunction] g && h [%v]\n", g && h) // -- conjunction --
	fmt.Printf("\n[Disjunction] g && h [%v]\n", g || h) // -- disjunction --

	// ----
}
