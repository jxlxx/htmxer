package htmxer

import (
	"embed"
	"fmt"
	"io/fs"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"golang.org/x/exp/maps"
)

type Server struct {
	Users map[string]*User
}

type User struct {
	Name        string
	ID          string
	Age         int
	Description string
}

func New() *Server {
	return &Server{
		Users: map[string]*User{
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
	user, userExists := s.Users[id]
	templ.Handler(editUserPage(user, userExists, false)).ServeHTTP(w, r)
}

// (DELETE /users/{id})
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, id string) {
	delete(s.Users, id)
	templ.Handler(userDeleted()).ServeHTTP(w, r)
}

// (POST /users/{id})
func (s Server) PostUser(w http.ResponseWriter, r *http.Request) {
	user, err := parseUserForm(r)
	if err != nil {
		templ.Handler(userNotFound()).ServeHTTP(w, r) // TODO: this should be a different error (maybe)
		return
	}
	id := randomID()
	s.Users[id] = userFormToUser(user, id)
	templ.Handler(userView(s.Users[id])).ServeHTTP(w, r)
}

// (PUT /users/{id})
func (s Server) PutUser(w http.ResponseWriter, r *http.Request, id string) {
	user, err := parseUserForm(r)
	if err != nil {
		templ.Handler(userNotFound()).ServeHTTP(w, r) // TODO: this should be a different error (maybe)
		return
	}
	s.Users[id] = userFormToUser(user, id)
	templ.Handler(userView(s.Users[id])).ServeHTTP(w, r)
}

// (GET /users/{id}/view)
func (s Server) ViewUser(w http.ResponseWriter, r *http.Request, id string) {
	user, ok := s.Users[id]
	if !ok {
		templ.Handler(userNotFound()).ServeHTTP(w, r)
		return
	}
	templ.Handler(userView(user)).ServeHTTP(w, r)
}

// (GET /users/{id}/view)
func (s Server) NewUser(w http.ResponseWriter, r *http.Request) {
	templ.Handler(editUserPage(nil, false, true)).ServeHTTP(w, r)
}

func parseUserForm(r *http.Request) (*UserForm, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	u := &UserForm{
		Name:        r.Form.Get("name"),
		Description: r.Form.Get("description"),
		Age:         parseInt(r.Form.Get("age")),
	}
	return u, nil
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func randomID() string {
	return fmt.Sprintf("user-%3d", rand.Intn(1000))
}

func userFormToUser(u *UserForm, id string) *User {
	return &User{
		ID:          id,
		Name:        u.Name,
		Description: u.Description,
		Age:         u.Age,
	}
}
