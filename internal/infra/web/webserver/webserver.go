package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handlers struct {
	Path     string
	Method   string
	Function http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      []Handlers
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []Handlers{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, Handlers{
		Method:   method,
		Path:     path,
		Function: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		switch handler.Method {
		case "GET":
			s.Router.Get(handler.Path, handler.Function)
		case "POST":
			s.Router.Post(handler.Path, handler.Function)
		case "PATCH":
			s.Router.Patch(handler.Path, handler.Function)
		case "DELETE":
			s.Router.Delete(handler.Path, handler.Function)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
