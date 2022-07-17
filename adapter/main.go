package main

import (
	"./data"
	"./dataService"
)

func main() {

	adapter()

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