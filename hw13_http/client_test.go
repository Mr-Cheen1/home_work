package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	address := defaultAddress
	port := 8080
	httpMethod := "GET"
	resourcePath := "/test"
	logger := SimpleLogger{}

	client := NewClient(address, port, httpMethod, resourcePath, logger)

	if client.address != address ||
		client.port != port ||
		client.httpMethod != httpMethod ||
		client.resourcePath != resourcePath ||
		client.logger != logger {
		t.Errorf("NewClient did not initialize the client struct correctly")
	}
}

func TestClientStart(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test response")
	}))
	defer testServer.Close()

	mockLogger := &MockLogger{}

	urlParts := strings.Split(testServer.URL, ":")
	portStr := urlParts[len(urlParts)-1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		t.Fatalf("Failed to convert port from string to int: %v", err)
	}

	client := NewClient(defaultAddress, port, "GET", "/", mockLogger)
	client.Start()

	expectedLogPart := "Client received response from"
	foundExpectedLog := false
	for _, msg := range mockLogger.Messages {
		if strings.Contains(msg, expectedLogPart) {
			foundExpectedLog = true
			break
		}
	}

	if !foundExpectedLog {
		t.Errorf("Expected log message to contain %q, but it was not found in logged messages", expectedLogPart)
	}
}
