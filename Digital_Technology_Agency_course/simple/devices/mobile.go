package devices

type Mobile struct {
	Cpu
}

func (mb Mobile) Off() {
	println("[Off] выключение телефона")
}

func (mb Mobile) On() {
	println("[On] включение телефона")
}
