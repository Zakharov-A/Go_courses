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
	Core  int
	Hard  int
	Brand string
}

func main() {
	var Test Pc = Pc{}
	Test.Core = 8
	Test.Brand = "ADM"
	fmt.Printf("\nCore = [%d] \nHard = [%d] \nBrand = [%s] \n", Test.Core, Test.Hard, Test.Brand)
	print("\nFinish\n")

}
