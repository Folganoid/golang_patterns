package main

import (
	"fmt"
	"./pkg"
)

var (
	base = pkg.BasePc{}
	home = pkg.HomePc{
		Cpu: 4,
		GraphicCard: 1,
		Wrapper: base,
	}
	server = pkg.ServerPc{
		Cpu: 16,
		Mem: 256,
		Wrapper: base,
	}
)


func main() {
	fmt.Printf("Base Price:[%f]\n", base.GetPrice())
	fmt.Printf("Home Cpu:[%d] / Cards:[%d] / Price:[%f]\n", home.Cpu, home.GraphicCard, home.GetPrice())
	fmt.Printf("Server Cpu:[%d] / Mem:[%d] / Price:[%f]\n", server.Cpu, server.Mem, server.GetPrice())
}