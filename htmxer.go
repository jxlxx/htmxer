package htmxer

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

//go:embed all:static
var static embed.FS

func (s Server) Static() fs.FS {
	return static
}

func (s Server) HomePage(w http.ResponseWriter, r *http.Request) {
	templ.Handler(homePage()).ServeHTTP(w, r)
}
