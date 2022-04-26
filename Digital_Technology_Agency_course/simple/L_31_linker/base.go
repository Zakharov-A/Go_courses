package L_31_linker

type Component interface {
	Search(componentName string)
	GetName() string
}