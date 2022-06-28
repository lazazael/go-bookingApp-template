package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/lazazael/go-bookingApp-template/pkg/config"
	"github.com/lazazael/go-bookingApp-template/pkg/handlers"
	"github.com/lazazael/go-bookingApp-template/pkg/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig

const portNumber string = ":8000"

var session *scs.SessionManager

//main is the application function
func main() {

	//app.InProduction set false for development
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
