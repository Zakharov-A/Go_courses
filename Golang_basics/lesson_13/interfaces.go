package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Man %s running", h.Name)
}

func main() {
	interfaceValues()
}

func interfaceValues() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	if runner == nil {
		fmt.Println("Runner is nil")
	}
	// runner = int64()
	// runner.Run()

	var unnamedRunner *Human
	fmt.Printf(" Type:%T Value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("Type %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("Runner is nil")
	}

	namedRunner := &Human{Name: "Dread"}
	fmt.Printf("Type: %T Value: %#v\n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = runner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = int64(1)
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

}
