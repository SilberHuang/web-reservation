package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SilberHuang/web-reservation/internal/config"
	"github.com/SilberHuang/web-reservation/internal/handlers"
	"github.com/SilberHuang/web-reservation/internal/models"
	"github.com/SilberHuang/web-reservation/internal/render"
	"github.com/alexedwards/scs/v2"
)

const PortNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	gob.Register(models.Reservation{})
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", PortNumber))

	srv := http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
