package main

import "fmt"

// ---- Unbuffered channel ----

// var inChannel chan int

// func main() {
// 	inChannel = make(chan int)
// 	go func() {
// 		println("Starting a thread")
// 		inChannel <- 10

// 	}()
// 	fmt.Println(<-inChannel)

// }

// ----

// ---- buffered channel ----

var inChannel chan int

func main() {
	inChannel = make(chan int, 3)
	inChannel <- 11
	inChannel <- 15
	inChannel <- 33
	fmt.Println(<-inChannel)
	fmt.Println(<-inChannel)
	fmt.Println(<-inChannel)

	println("Starting a thread")

}

// ----
