package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Server struct {
	address    string
	port       int
	logger     Logger
	httpServer *http.Server
	db         DatabaseInterface
}

type DatabaseInterface interface {
	ExecuteSQLFromFile(filePath string) (interface{}, error)
}

func NewServer(address string, port int, logger Logger, db DatabaseInterface) *Server {
	return &Server{
		address: address,
		port:    port,
		logger:  logger,
		db:      db,
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

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	if s.logger == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/execute-sql" {
		s.logger.Log(fmt.Sprintf("Path %s not found", r.URL.Path))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Обработка GET запроса.
		s.logger.Log("Received a GET request at /execute-sql")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("This is a response to your GET request at /execute-sql")); err != nil {
			s.logger.Log(fmt.Sprintf("Error writing response: %v", err))
		}
	case http.MethodPost:
		// Обработка POST запроса.
		var requestData struct {
			SQLFilePath string `json:"sqlFilePath"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		defer r.Body.Close()
		if err != nil {
			s.logger.Log(fmt.Sprintf("Error decoding request body: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fileName := requestData.SQLFilePath
		result, err := s.db.ExecuteSQLFromFile("sql_scripts/" + fileName)
		if err != nil {
			s.logger.Log(fmt.Sprintf("Error executing SQL script: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			s.logger.Log(fmt.Sprintf("Error writing response: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		s.logger.Log(fmt.Sprintf("Method %s not allowed for /execute-sql", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
