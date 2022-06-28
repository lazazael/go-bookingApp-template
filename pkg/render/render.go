package render

import (
	"bytes"
	"github.com/lazazael/go-bookingApp-template/pkg/config"
	"github.com/lazazael/go-bookingApp-template/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//functions is not built into the templates_functions
var functions = template.FuncMap{}

var app *config.AppConfig

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

//AddDefaultData add data here which is available for every page
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td

}

//RenderTemplate renders the templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache:")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("error wringing template to browser", err)
	}
}

//CreateTemplateCache creates a hashmap as cache of available templates from the templates directory
func CreateTemplateCache() (map[string]*template.Template, error) {

	myTemplateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myTemplateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//ts is a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myTemplateCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myTemplateCache, err
		}
		//
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myTemplateCache, err
			}
		}

		myTemplateCache[name] = ts
	}
	return myTemplateCache, nil
}
