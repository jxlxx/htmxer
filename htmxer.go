package htmxer

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"golang.org/x/exp/maps"
)

type Server struct {
	Users map[string]User
}

type User struct {
	Name        string
	ID          string
	Age         int
	Description string
}

func New() *Server {
	return &Server{
		Users: map[string]User{
			"user-1": {
				ID:          "user-1",
				Name:        "Alpha",
				Age:         25,
				Description: "random description",
			},
			"user-2": {
				ID:          "user-2",
				Name:        "Bravo",
				Age:         25,
				Description: "random description",
			},
			"user-3": {
				ID:          "user-3",
				Name:        "Charlie",
				Age:         25,
				Description: "random description",
			},
		},
	}
}

//go:embed all:static
var static embed.FS

func (s Server) Static() fs.FS {
	return static
}

// (GET /)
func (s Server) HomePage(w http.ResponseWriter, r *http.Request) {
	templ.Handler(homePage()).ServeHTTP(w, r)
}

// (GET users/)
func (s Server) UserListPage(w http.ResponseWriter, r *http.Request) {
	users := maps.Values(s.Users)
	templ.Handler(usersListPage(users)).ServeHTTP(w, r)
}

// (GET /users/{id})
func (s Server) UserPage(w http.ResponseWriter, r *http.Request, id string) {
	user, ok := s.Users[id]
	templ.Handler(userPage(user, ok)).ServeHTTP(w, r)
}

// (GET /users/{id}/edit)
func (s Server) EditUser(w http.ResponseWriter, r *http.Request, id string) {
	user, ok := s.Users[id]
	templ.Handler(editUserPage(user, ok)).ServeHTTP(w, r)
}

// (DELETE /users/{id})
func (s Server) DeleteUser(w http.ResponseWriter, r *http.Request, id string) {

}

// (POST /users/{id})
func (s Server) PostUser(w http.ResponseWriter, r *http.Request) {

}

// (PUT /users/{id})
func (s Server) PutUser(w http.ResponseWriter, r *http.Request, id string) {

}
