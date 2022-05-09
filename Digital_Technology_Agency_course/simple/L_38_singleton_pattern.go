package main

import L_38 "simple/L_38"

var singleton *L_38.Singleton

func init() {
	println("Basic program initialization")
	singleton = &L_38.Singleton{"Singleton"}
}

func main() {
	for i := 0; i < 3; i++ {
		singleton = L_38.NewSingleton(singleton, "Create singleton")
	}

}
