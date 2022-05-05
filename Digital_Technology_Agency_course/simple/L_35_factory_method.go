package main

import "simple/L_35"


var types = []string{L_35.PersonalComputerType, L_35.NotebookType, L_35.ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := L_35.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}