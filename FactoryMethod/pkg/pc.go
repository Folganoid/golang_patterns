package pkg

import "fmt"

type PC struct {
	Type string
	Core int
	Memory int
	Monitor bool
}

func NewPC() Computer {
	return PC{
		Type: PCType,
		Core: 4,
		Memory: 32,
		Monitor: false,
	}
}

func (pc PC) GetType() string {
	return pc.Type
}

func (pc PC) PrintDetails() {
	fmt.Printf("%s Core:[%d], Mem: [%d], Monitor: [%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
