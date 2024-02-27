package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	address    string
	port       int
	logger     Logger
	httpServer *http.Server
}

func NewServer(address string, port int, logger Logger) *Server {
	return &Server{
		address: address,
		port:    port,
		logger:  logger,
		httpServer: &http.Server{
			Addr:              fmt.Sprintf("%s:%d", address, port),
			ReadHeaderTimeout: 15 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	if s.logger == nil {
		return fmt.Errorf("logger is not initialized")
	}

	s.httpServer.Handler = http.HandlerFunc(s.handleRequest)
	s.logger.Log(fmt.Sprintf("Attempting to start server on %s", s.httpServer.Addr))

	ln, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		errorMsg := fmt.Sprintf("Error starting server: %v", err)
		if strings.Contains(err.Error(), "bind: address already in use") ||
			strings.Contains(err.Error(), "listen tcp :8080: bind: An attempt was made to access a socket") {
			errorMsg = "Error starting server: Port is already in use"
		}
		s.logger.Log(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	go func() {
		defer ln.Close()
		if err = s.httpServer.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Log(fmt.Sprintf("Server error: %v", err))
		}
	}()

	return nil
}

func (s *Server) Stop() {
	if s.logger == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.logger.Log("Attempting to stop server")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Log(fmt.Sprintf("Server shutdown error: %v", err))
	} else {
		s.logger.Log("Server successfully stopped")
	}
}

func (s Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	if s.logger == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.logger.Log(fmt.Sprintf("Received %s request to %s", r.Method, r.URL.Path))

	// Сервер принимает HTTP GET и POST запросы и возвращает ответ с данными клиенту.
	switch r.Method {
	case "GET":
		s.logger.Log("Processing GET request")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("GET request processed"))
		if err != nil {
			s.logger.Log(fmt.Sprintf("Error writing GET response: %v", err))
		}
	case "POST":
		s.logger.Log("Processing POST request")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("POST request processed"))
		if err != nil {
			s.logger.Log(fmt.Sprintf("Error writing POST response: %v", err))
		}
	default:
		s.logger.Log(fmt.Sprintf("Unsupported method: %s", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			s.logger.Log(fmt.Sprintf("Error writing Method Not Allowed response: %v", err))
		}
	}
}
