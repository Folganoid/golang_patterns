package pkg

import "fmt"

type Cargo struct {
	Dispatcher
}

func (g *Cargo) PermitArrive() {
	fmt.Println("Cargo: loading.")
	g.Arrive()
}

func (g *Cargo) Go() {
	fmt.Println("Cargo: loaded.")
	g.NotifyAboutGo()
}

func (g *Cargo) Arrive() {
	if !g.CanArrive(g) {
		fmt.Println("Cargo: loading is interrupt. Wait.")
		return
	}
	fmt.Println("Cargo: go!!!.")
}