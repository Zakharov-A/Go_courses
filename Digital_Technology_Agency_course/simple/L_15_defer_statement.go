package main

import "fmt"

// ---- defer ----

func main() {

	fmt.Println("main begin")
	defer end()
	defer begin()
	defer red()
}

func begin() {
	fmt.Println("Hello my friends")
}

func end() {
	fmt.Println("Word at the end")
}

func red() {
	fmt.Println("I'm red, I'm best!")
}

// ----
