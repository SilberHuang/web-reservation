package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
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
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

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

func run() error {
	gob.Register(models.Reservation{})
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	return nil
}
