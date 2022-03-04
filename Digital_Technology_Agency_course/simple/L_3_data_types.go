package main

import "fmt"

// ---- Data types ----

var i8 int8 = 15
var i16 int16 = -20345
var u64 uint64 = 50965
var ru5 rune = -214345
var fl32 float32 = 3.4 * 10.0
var comp64 complex64 = 3.4 * 18.8

func main() {
	fmt.Printf(`%f`, comp64)

}

func Run(print string) {
	fmt.Printf(`#{print}`)
}

// ----
