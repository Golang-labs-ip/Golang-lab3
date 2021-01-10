package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Golang-labs-ip/Golang-lab3/server/tablets"
)

type HttpPortNumber int

type TabletApiServer struct {
	Port HttpPortNumber

	TabletsHandler tablets.HTTPHandlerFunc

	server *http.Server
}

func (s *TabletApiServer) Start() error {
	if s.TabletsHandler == nil {
		return fmt.Errorf("tablets HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/tablets", s.TabletsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *TabletApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
