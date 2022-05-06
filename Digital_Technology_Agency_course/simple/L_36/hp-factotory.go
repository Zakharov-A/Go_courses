package L_36

type HpFactory struct {
}

func (factory HpFactory) GetComputer() Computer {
	return HpComputer{
		Memory: 8,
		Cpu: 4,
	}
}

func (factory HpFactory) GetMonitor() Monitor {
	return HpMonitor{
		Size: 24,
	}
}