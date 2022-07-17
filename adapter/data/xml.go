package data

type XmlDocument struct {
	body string
}

func (doc *XmlDocument) SetBodyData(body string) {
	doc.body = body
}

func (doc XmlDocument) GetBodyData() string {
	return doc.body
}

func (doc XmlDocument) SendXmlData() {
	println("Send XML: " + doc.GetBodyData())
}
