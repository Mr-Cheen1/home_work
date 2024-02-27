package main

import (
	"flag"
	"os"
	"reflect"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestParseFlags(t *testing.T) {
	testCases := []struct {
		name        string
		args        []string
		expected    Config
		expectError bool
	}{
		{
			name: "Server mode with full options",
			args: []string{
				"-mode", "server",
				"-address", "127.0.0.1",
				"-port", "8080",
				"-method", "get",
				"-path", "/test",
			},
			expected: Config{
				mode:         "server",
				address:      "127.0.0.1",
				port:         8080,
				httpMethod:   "get",
				resourcePath: "/test",
			},
			expectError: false,
		},
		{
			name: "Client mode with full options",
			args: []string{
				"-mode", "client",
				"-address", "192.168.1.1",
				"-port", "9090",
				"-method", "post",
				"-path", "/data",
			},
			expected: Config{
				mode:         "client",
				address:      "192.168.1.1",
				port:         9090,
				httpMethod:   "post",
				resourcePath: "/data",
			},
			expectError: false,
		},
		{
			name: "Server mode with default options",
			args: []string{
				"-mode", "server",
			},
			expected: Config{
				mode:         "server",
				address:      defaultAddress,
				port:         8080,
				httpMethod:   "get",
				resourcePath: "/",
			},
			expectError: false,
		},
		{
			name: "Default mode and options",
			args: []string{},
			expected: Config{
				mode:         "server",
				address:      defaultAddress,
				port:         8080,
				httpMethod:   "get",
				resourcePath: "/",
			},
			expectError: false,
		},
		{
			name: "Invalid port",
			args: []string{
				"-mode", "client",
				"-port", "wrong",
			},
			expected:    Config{},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Сохраняем оригинальные аргументы командной строки и флаги.
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()
			os.Args = append([]string{"program"}, tc.args...)

			// Сбрасываем флаги перед каждым тестом.
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			cfg, err := parseFlags()

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error for args %v, but no error was returned", tc.args)
				}
			} else {
				if err != nil {
					t.Errorf("Returned an unexpected error for args %v: %v", tc.args, err)
				} else if !reflect.DeepEqual(*cfg, tc.expected) {
					t.Errorf("parseFlags() = %+v, want %+v for args %v", *cfg, tc.expected, tc.args)
				}
			}
		})
	}
}

// MockServer имитирует сервер для тестирования.
type MockServer struct {
	Stopped bool
	mu      sync.Mutex
}

func (m *MockServer) Stop() {
	m.mu.Lock()
	m.Stopped = true
	m.mu.Unlock()
}

func TestWaitForShutdown(t *testing.T) {
	mockServer := &MockServer{}
	sigs := make(chan os.Signal, 1)

	go waitForShutdown(mockServer, sigs)

	sigs <- syscall.SIGTERM

	time.Sleep(100 * time.Millisecond)

	mockServer.mu.Lock()
	if !mockServer.Stopped {
		mockServer.mu.Unlock() // Не забываем разблокировать мьютекс
		t.Errorf("Server was not stopped by waitForShutdown")
	}
	mockServer.mu.Unlock()
}
