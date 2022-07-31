package main

import (
	"fmt"
	"./pkg"
)

var routes = pkg.Routes{
	Routes: []pkg.Route{
		{Name: "Route-1", TravelTime: 110},
		{Name: "Route-2", TravelTime: 50},
		{Name: "Route-3", TravelTime: 60},
		{Name: "Route-4", TravelTime: 40},
		{Name: "Route-5", TravelTime: 35},
	},
}


func main() {
	for routes.HasNext() {
		route := routes.GetNext()
		fmt.Printf("R:[%s] Time[%d]\n", route.Name, route.TravelTime)
	}
}