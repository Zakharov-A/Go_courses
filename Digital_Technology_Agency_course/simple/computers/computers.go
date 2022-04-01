package devices

type Cpu struct {
	Prod string
	Core int
	Hard int
}

type Pc struct {
	Cpu
}

func (pc Pc) On() {
	println("[On] Shutdown PC")
}

func (pc Pc) Off() {
	println("[Off] Turning on PC")
}

func (pc Pc) initMethod() {
	println("[Init] ... ")
}
