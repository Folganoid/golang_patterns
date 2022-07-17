package data

type JsonDocument struct {
	body string
}

type JsonDocumentAdapter struct {
	JsonDocument *JsonDocument
}

func (doc *JsonDocument) SetBodyData(body string) {
	doc.body = body
}

func (adapter JsonDocumentAdapter) GetBodyData() string {
	return adapter.JsonDocument.body
}

func (doc *JsonDocument) сonvertJsonToXml() {
	doc.body = "<xml>" + doc.body + "</xml>"
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.JsonDocument.сonvertJsonToXml()
	print("Send JSON->XML data: " + adapter.JsonDocument.body)
}
