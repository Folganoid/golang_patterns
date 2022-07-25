package main

import "./pkg"

var singleton *pkg.Singleton

func init() {
	println("Base init")
	singleton = &pkg.Singleton{Type: "Odino4ka"}
}

func main() {
	for i := 0 ; i < 3; i++ {
		pkg.NewSingleton(singleton, "Odinochka+++" )
	}
}
