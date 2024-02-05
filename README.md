# Simple HTMX/Tailwind/Go example

This module is a simple example of a web server using [`templ`](https://templ.guide/) + [`htmx`](https://htmx.org/) + [`tailwind`](https://tailwindcss.com/).

The web server maintains a map of `User`s and the frontend allows vistors to
the site to view, create, edit, and destroy `Users`. 

It is extremely basic and only intended to demonstrate some basic elements of the
featured tools.

## Features

- :sparkles: CRUD actions implemented with [htmx](https://htmx.org/)
- :sparkles: HTML elements generated with [templ](https://github.com/a-h/templ)
- :sparkles: styled with [Tailwind](https://tailwindcss.com/)
- :sparkles: non-tailwind [branch](https://github.com/jxlxx/htmxer/tree/no-tailwind) available 
- :sparkles: runs with [Task](https://taskfile.dev/)
- :sparkles: live reloads with [Air](https://github.com/cosmtrek/air)
- :sparkles: tailwind and htmx are vendor'ed
- :sparkles: server interface generated with [oapi-codegen](https://github.com/deepmap/oapi-codegen)

## How to run

#### if you have `task` and `air` installed

> (you need `task` because `air` uses a `task` command to build the binary it runs).

```sh
air
```

#### if you don't have task or air installed (no live reload)

```sh
go run cmd/main.go
```

#### installing  `task` and `air` 

Assuming you already have Go 1.21+ installed, you can just use the `go install` command.

```sh
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/cosmtrek/air@latest
```

But there are lots of ways to install both.

## How to develop


You don't really *need* `air` and `task`, but they are helpful.
```sh
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/cosmtrek/air@latest
```

But you will need `oapi-codegen` + `templ`.

```sh
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
go install github.com/a-h/templ/cmd/templ@latest
```

### New routes

1. Add a path to `openapi.yaml`

```yaml
paths:
  /:
    get: 
      operationId: HomePage
  # ...
  # ...
  # ...
  /new/route:
    get:
      operationId: NewRoute
```

2. Generate the code with `task` (or look at the relevant commands in the `Taskfile.yml`).

```sh
task api
```

3. Create a whatever html you want with `templ` components.

```go
templ newRoute() {
  @baseLayout("new route") {
    <div>
      This is a new page.
    </div>
  }
}
```

4. Implement the interface as method to the server and render the html.

```go
func (s Server) NewRoute() {
  templ.Handler(newRoutePage()).ServeHTTP(w, r)
}
```

5. Run the server with `air` and go to [`localhost:800/new/route`](http://localhost:8000/new/route)

```sh
air
```
