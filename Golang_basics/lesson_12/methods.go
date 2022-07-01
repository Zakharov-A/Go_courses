package main

import "fmt"

// type OurType string

// 1. We can do this because the type is alredy in this package.
// func (t OurType) Hello() {
// 	fmt.Println("Hello girls!")
// }

func main() {

	definition()

}

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("figure perimetter: %d \n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier

}

func definition() {
	square := Square{Side: 4}
	pSquare := &square

	square2 := Square{Side: 2}

	square.Perimeter()
	square2.Perimeter()

	pSquare.Scale(2)
}
