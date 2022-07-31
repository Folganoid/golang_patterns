package main

import "./pkg"

func main() {

	device := &pkg.Device{Name: "Device1"}
	updateSvc := &pkg.UpdateDataService{Name: "Update1"}
	dataSvc := &pkg.DataService{}
	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)
	data := &pkg.Data{}
	device.Execute(data)

}