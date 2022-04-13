package data

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToxml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	return adapter.jsonDocument.ConvertToxml()
	println("Sending Xml data!")
}
