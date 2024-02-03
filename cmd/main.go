package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/jxlxx/htmxer"
)

func main() {
	r := chi.NewRouter()
	handler := htmxer.New()

	r.Mount("/", htmxer.Handler(handler))

	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.FS(handler.Static()))
		fs.ServeHTTP(w, r)
	})

	fmt.Println("listening on http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", r); err != nil {
		fmt.Println(err)
	}
}
