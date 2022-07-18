package pkg

import "fmt"

type Cpu struct {
	Name string
	Description string
}

func (cpu Cpu) GetName() string {
	return cpu.Name
}

func (cpu Cpu) Search(name string) {
	if cpu.Name == name {
		fmt.Printf("Component [%s] found %s", cpu.Name, cpu.Description)
	}
}