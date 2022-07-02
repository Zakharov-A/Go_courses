package main

import (
	"fmt"
	"time"
)

type OurType string

// 1. We can do this because the type is alredy in this package.
func (t OurType) Hello() {
	fmt.Println("Hello girls!")
}

type plnt *int

// func (plnt) Test() {
// 	fmt.Println("TEST")
// }

type Tester interface {
	Hello()
}

// func (t Tester) Test() {
// 	fmt.Println("Test")
// }

// func (i int64) Test() {
// 	fmt.Println("Test")
// }

func main() {

	// definition()
	rules()

}

// type Square struct {
// 	Side int
// }

// func (s Square) Perimeter() {
// 	fmt.Printf("%T, %#v \n", s, s)
// 	fmt.Printf("figure perimetter: %d \n", s.Side*4)
// }

// func (s *Square) Scale(multiplier int) {
// 	fmt.Printf("%T, %#v \n", s, s)
// 	s.Side *= multiplier
// 	fmt.Printf("%T, %#v \n", s, s)

// }

// func (s Square) WrongScale(multiplier int) {
// 	fmt.Printf("%T, %#v, \n", s, s)
// 	s.Side *= multiplier
// 	fmt.Printf("%T, %#v \n", s, s)
// }

// func definition() {
// 	square := Square{Side: 4}
// 	pSquare := &square

// 	square2 := Square{Side: 2}

// 	square.Perimeter()
// 	square2.Perimeter()

// 	pSquare.Scale(2)

// 	pSquare.Perimeter()
// 	square.Scale(2)
// 	pSquare.Perimeter()

// 	square.WrongScale(2)
// 	square.Perimeter()
// }

func rules() {
	// 1. The type must be declared in the curret pacage.
	// 5. Is a declared type

	now := time.Now()
	fmt.Printf("%T, %#v \n", now, now)

	var ourString OurType = "Hello"
	fmt.Printf("%T, %#v \n", ourString, ourString)

	// 2. Type must not be a pointer.
	var plnt plnt
	fmt.Print("T%, %#v \n", plnt, plnt)

	// 3. Interface type
	var stringTwo Tester
	fmt.Printf("%T, %#v \n", stringTwo, stringTwo)

	// 4. Builtin type
	var builtinType int64 = 1
	fmt.Printf("%T, %#v \n", builtinType, builtinType)

}
