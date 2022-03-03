package mySite

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Millisecond,
		WriteTimeout:   10 * time.Millisecond,
	}
	return s.httpServer.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) error {

	return s.httpServer.Shutdown(ctx)

}
