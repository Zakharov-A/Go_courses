package main

import (
	"fmt"
	"simple/L_40"
)

func main() {
	storage := &L_40.Guardian{
		Items: make([]*L_40.Snapshot, 0),
	}
	creator := &L_40.Creator{
		State: "A",
	}
	fmt.Printf("Creator current State: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("B")
	fmt.Printf("Creator current State: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("C")
	fmt.Printf("Creator current State: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.Restore(storage.Get(1))
	fmt.Printf("Restored to state: %s\n", creator.GetState())

	creator.Restore(storage.Get(0))
	fmt.Printf("Restored to state: %s\n", creator.GetState())

}
