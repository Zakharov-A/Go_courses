package main

import "fmt"

var i8 int8 = 15
var i16 int16 = -20345
var u64 uint64 = 50965
var ru5 rune = -214345
var fl32 float32 = 3.4 * 10.0

func main() {
	fmt.Printf(`%f`, fl32)

}

func Run(print string) {
	fmt.Printf(`#{print}`)
}
