package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type Server struct {
	// TODO add server config here
	httpServer    http.Server
	syncWaitGroup sync.WaitGroup
}

func New() (*Server, error) {
	// TODO New with server config
	// TODO allow custom engine - currently gin.Default()
	e := gin.Default()
	return &Server{
		httpServer:    http.Server{Addr: ":8080", Handler: e},
		syncWaitGroup: sync.WaitGroup{},
	}, nil
}

// ListenAndServe - see also http.Server.ListenAndServe (https://pkg.go.dev/net/http#Server.ListenAndServe)
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown - see also http.Server.Shutdown (https://pkg.go.dev/net/http#Server.Shutdown)
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
