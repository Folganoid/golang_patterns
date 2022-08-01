package pkg

import "fmt"

type Passenger struct {
	Dispatcher
}

func (g *Passenger) PermitArrive() {
	fmt.Println("Passengers: ready to go.")
	g.Arrive()
}

func (g *Passenger) Go() {
	fmt.Println("Passengers: are going.")
	g.NotifyAboutGo()
}

func (g *Passenger) Arrive() {
	if !g.CanArrive(g) {
		fmt.Println("Passengers: going is interrupt. Wait.")
		return
	}
	fmt.Println("Passengers: ready to go.")
}

