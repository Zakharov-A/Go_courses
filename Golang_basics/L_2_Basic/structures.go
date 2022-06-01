package main

import "fmt"

// ----

// type Pc struct {
// 	Core  int
// 	Hard  int
// 	Brand string
// }

// func main() {
// 	var Test Pc = Pc{
// 		Core:  4,
// 		Hard:  250,
// 		Brand: "Intel",
// 	}
// 	fmt.Printf("\nCore = [%d] \nHard = [%d] \nBrand = [%s]", Test.Core, Test.Hard, Test.Brand)

// }

// ----

type Pc struct {
	Core   int
	Hard   int
	Brand  string
	Router string
}

func main() {
	var Test Pc = Pc{}
	Test.Hard = 1000
	Test.Core = 8
	Test.Brand = "ADM"
	Test.Router = "Mi"
	fmt.Printf("\nCore = [%d] \nHard = [%d] \nBrand = [%s] \nRouter = [%s]\n", Test.Core, Test.Hard, Test.Brand, Test.Router)

}
