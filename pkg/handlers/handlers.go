package handlers

import (
	"net/http"

	"github.com/SilberHuang/web-reservation/pkg/config"
	"github.com/SilberHuang/web-reservation/pkg/models"
	"github.com/SilberHuang/web-reservation/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
