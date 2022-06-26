package main

import "fmt"

func main() {
	// default value
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)

	var intPointer2 *string
	fmt.Printf("%T %#v \n", intPointer2, intPointer2)

	// Getting not-nil pointers
	// variadle
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	var pointerA = &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

	// get pointer via new keyword
	var newPoiter = new(float32)
	fmt.Printf("%T %#v %#v \n", newPoiter, newPoiter, *newPoiter)
	*newPoiter = 3
	fmt.Printf("%T %#v %#v \n", newPoiter, newPoiter, *newPoiter)

}
