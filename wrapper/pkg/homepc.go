package pkg

type HomePc struct {
	Cpu int
	GraphicCard int
	Wrapper Wrapper

}

func (pc HomePc) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.Cpu) + pc.Wrapper.GetPrice() * float64(pc.GraphicCard)
}