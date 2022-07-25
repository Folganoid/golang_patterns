package pkg

type HpCollector struct {
	Core int
	Brand string
	Memory int
	Monitor int
	GraphicCard int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 8
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 32
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 3
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 3
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer {
		Core: collector.Core,
		Brand: collector.Brand,
		Memory: collector.Memory,
		Monitor: collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}