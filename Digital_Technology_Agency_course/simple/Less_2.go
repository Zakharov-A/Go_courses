package main

import "fmt"

// var First = "Good Word" //Глобал

// var first = "word" // Переменная с маленькой буквы внутри пакета

func main() {
	first := "Roghan the best!"
	Run(first)

}

func Run(xxx string) {
	fmt.Printf("%s", xxx)
}
