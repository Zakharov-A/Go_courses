package main

import "time"

// ---- Goroutines. First acquaintance with parallel programming ----

func main() {
	simple()
	parallel()

}

func simple() {
	var arr = []int{1, 2, 3, 4, 5}
	for _, item := range arr {
		println(item)
	}

}

func parallel() {
	var arr = []int{1, 2, 3, 4, 5}
	for _, item := range arr {
		go println(item)
	}
	time.Sleep(100 * time.Millisecond)
	print("Method completed")
}

// ----
