package dataService

type AnalyticalDataService interface {
	SendXmlData()
	GetBodyData() string
}

func SendXml(service AnalyticalDataService) {
	service.SendXmlData()
}

func GetBody(service AnalyticalDataService) string {
	return service.GetBodyData()
}
