package L_31_linker

import "fmt"

type Motherboard struct {
	Name string
	Description string
	Components []Component
}

func (mb Motherboard) GetName() string{
	return mb.Name
}

func (mb Motherboard) Search(name string) {
	if mb.Name == name {
		fmt.Printf("Component [%s] found %s", mb.Name, mb.Description)
		return
	}

	for _, component := range mb.Components {
		component.Search(name)
	}
}