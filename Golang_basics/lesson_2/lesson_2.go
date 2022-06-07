package main

import (
	"fmt"
	"unsafe"
)

// ----
// func main() {
// 	name := "Sam"
// 	hello := fmt.Sprintf("hello %s", name)
// 	// _ = hello
// 	fmt.Println(hello)

// 	fmt.Printf("Type: %T Value: %v\n", name, name)

// }
// ----

// var hi string = "Hi you"
// var hello string = "Fack you!"

// const name = "Sam"

// func main() {

// 	// name := "John"
// 	fmt.Println(name)

// var hello string = "Hi my friend!"

// hi := "Cool"
// ourBool := true
// fmt.Println(hello)
// fmt.Println(ourBool)
// fmt.Println(hi)

// outSting := fmt.Sprintf("our string: %T", hello)
// fmt.Println(outSting)

// fmt.Printf("Type: %T Value: %v\n", hello, hello)
// 	fmt.Printf("Type: %T Value: %v\n", name, name)
// 	hihi()

// }

// func hihi() {
// 	fmt.Println(hello)
// }

// ----

func main() {
	var num1 uint8 = 1
	var num2 uint64 = 1

	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))

}
