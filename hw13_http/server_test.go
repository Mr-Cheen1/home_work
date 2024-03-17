package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MockLogger struct {
	LogCalls int
	Messages []string
}

func (l *MockLogger) Log(message string) {
	l.LogCalls++
	l.Messages = append(l.Messages, message)
}

func TestNewServer(t *testing.T) {
	logger := MockLogger{}
	server := NewServer(defaultAddress, 8080, &logger)

	if server.address != defaultAddress || server.port != 8080 {
		t.Errorf("NewServer did not correctly assign address or port")
	}
}

func TestHandleRequest(t *testing.T) {
	logger := MockLogger{}
	server := NewServer(defaultAddress, 8080, &logger)
	server.httpServer.Handler = http.HandlerFunc(server.handleRequest)

	ts := httptest.NewServer(server.httpServer.Handler)
	defer ts.Close()

	testPostBody := "test body"
	expectedPostResponseBody := "POST request processed: /test, with body: " + testPostBody

	tests := []struct {
		method       string
		want         int
		body         string
		expectedBody string
	}{
		{"GET", http.StatusOK, "", "GET request processed: /test"},
		{"POST", http.StatusOK, testPostBody, expectedPostResponseBody},
		{"PUT", http.StatusMethodNotAllowed, "", ""},
	}

	initialLogCalls := logger.LogCalls

	for _, tc := range tests {
		var bodyReader io.Reader
		if tc.body != "" {
			bodyReader = strings.NewReader(tc.body)
		}

		req, err := http.NewRequestWithContext(
			context.Background(),
			tc.method,
			ts.URL+"/test",
			bodyReader,
		)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != tc.want {
			t.Errorf("For method %s expected status %d, got %d", tc.method, tc.want, resp.StatusCode)
		}

		// Проверка тела ответа.
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != tc.expectedBody && tc.method != "PUT" {
			t.Errorf("For method %s expected body %s, got %s", tc.method, tc.expectedBody, string(body))
		}
	}
	if logger.LogCalls <= initialLogCalls {
		t.Errorf("Expected Log to be called during request handling, but it was not")
	}
}

func TestServerStartAndStop(t *testing.T) {
	logger := MockLogger{}
	server := NewServer(defaultAddress, 8081, &logger)

	initialLogCalls := logger.LogCalls

	errChan := make(chan error, 1)
	go func() {
		err := server.Start()
		errChan <- err
	}()

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatalf("Server failed to start: %v", err)
		}
	case <-time.After(1 * time.Second):
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://localhost:%d/", 8081), nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make GET request to the server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	server.Stop()

	time.Sleep(1 * time.Second)

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err = http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://localhost:%d/", 8081), nil)
	if err != nil {
		t.Fatalf("Failed to create request after server stop: %v", err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err == nil {
		resp.Body.Close()
		t.Fatalf("Server did not shut down properly")
	}

	if logger.LogCalls == initialLogCalls {
		t.Errorf("Log on server start/stop is not called")
	}
}

func TestServerStartOnOccupiedPort(t *testing.T) {
	logger := MockLogger{}
	server1 := NewServer(defaultAddress, 8082, &logger)

	if err := server1.Start(); err != nil {
		t.Fatalf("Failed to start the first server: %v", err)
	}
	defer server1.Stop()

	time.Sleep(1 * time.Second)

	initialLogCalls := logger.LogCalls

	server2 := NewServer(defaultAddress, 8082, &logger)

	err := server2.Start()
	if err == nil {
		t.Errorf("Expected error when starting second server on occupied port, got nil")
	}
	defer server2.Stop()

	if logger.LogCalls == initialLogCalls {
		t.Errorf("Expected Log to be called when starting second server on occupied port, but it was not")
	}
}
