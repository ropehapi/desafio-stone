package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) RegisterRoutes(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Serve() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}
	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		panic(err)
	}
}
