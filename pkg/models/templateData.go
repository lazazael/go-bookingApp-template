package models

//TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any //interface{}
	CSRFToken string         //Cross-Site Request Forgery Token (security token for handling tokens)

	Flash   string
	Warning string
	Error   string
}
