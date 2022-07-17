package main

import (
	"./types"

	"./data"
	"./dataService"
)

func main() {

	bridge()
	//adapter()

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


func adapter() {
	//send xml
	xml := data.XmlDocument{}
	xml.SetBodyData("<xml>111</xml>")
	dataService.SendXml(xml)
	dataService.GetBody(xml)

	//send converted json
	var json = data.JsonDocument{}
	json.SetBodyData(`{"aaa": 111, "bbb": 222}`)
	adapt := data.JsonDocumentAdapter{}
	adapt.JsonDocument = &json
	dataService.SendXml(adapt)
	dataService.GetBody(adapt)
}