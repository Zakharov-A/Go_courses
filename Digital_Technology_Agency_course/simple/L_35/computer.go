package L_35

import "fmt"

const (
	ServerType = "server"
	PersonalComputerType = "personal"
	NotebookType = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s Not an existing object type!\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()				
	}
}