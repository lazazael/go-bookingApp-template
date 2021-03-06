package handlers

import (
	"github.com/lazazael/go-bookingApp-template/pkg/config"
	"github.com/lazazael/go-bookingApp-template/pkg/models"
	"github.com/lazazael/go-bookingApp-template/pkg/render"
	"net/http"
)

//Repository pattern allows swapping components around the application

//Repository the repository type
type Repository struct {
	App *config.AppConfig
}

//Repo the repository used by the handlers
var Repo *Repository

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//_, _= fmt.Fprintf(w,"This is the home page.")

	//storing the remoteIP on call
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform logic
	stringMap := make(map[string]string)
	stringMap["test"] = "test string on about page"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
