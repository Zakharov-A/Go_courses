package devices

type Pc struct {
	Cpu
}

func (pc Pc) Off() {
	println("[off] выключение PC")
}

func (pc Pc) On() {
	println("[On] Включение PC")
}
