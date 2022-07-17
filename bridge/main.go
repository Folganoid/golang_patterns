package main

import (
	"./types"
)

func main() {

	bridge()

}

func bridge() {

	hpScanner := types.HP{}
	epsonScanner := types.Epson{}
	
	linuxPc := types.LinuxPC{}
	winPc := types.WinPC{}
	//macPc := types.MacPC{}


	linuxPc.AddScanner(hpScanner)
	winPc.AddScanner(hpScanner)
	linuxPc.Scan()
	winPc.Scan()
	linuxPc.AddScanner(epsonScanner)
	linuxPc.Scan()

}