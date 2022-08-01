package main

import (
	"time"

	"./pkg"
)

func main() {
	stationManager := pkg.NewStationManager()
	passengerBus := &pkg.Passenger{
		Dispatcher: stationManager,
	}

	cargo := &pkg.Cargo{
		Dispatcher: stationManager,
	}

	passengerBus.Arrive()
	time.Sleep(time.Second * 1)
	cargo.Arrive()
	time.Sleep(time.Second * 1)
	passengerBus.Go()
}