package pkg

type HpFactory struct {

}

func (factory HpFactory) GetComputer() Computer {
	return HpComputer{
		Memory: 8,
		Cpu: 2,
	}
}

func (factory HpFactory) GetMonitor() Monitor {
	return HpMonitor{
		Size: 15,
	}
}
