package main

import (
	"os"
)

// ---- Requesting information and output to the console ----

// func main() {
// 	var name string
// 	println("what is your name?")
// 	_, err := fmt.Scan(&name)
// 	if err != nil {
// 		println("Console input error")
// 		return
// 	}
// 	fmt.Println("Your name:", name)
// }

// ----

// ---- Cat utility ----

func main() {
	args := os.Args
	if len(args) < 2 {
		println("Need to pass arguments!")
		return
	}
	for i := 1; i < len(args); i++ {
		fileData, err := os.ReadFile(args[i])
		if err != nil {
			println("Error:", err.Error())
			return
		}
		println(string(fileData))
	}

}

// ----
