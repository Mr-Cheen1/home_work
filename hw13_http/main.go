package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const defaultAddress = "localhost"

type Config struct {
	mode         string
	address      string
	port         int
	httpMethod   string
	resourcePath string
}

type Logger interface {
	Log(message string)
}

type SimpleLogger struct{}

func (l SimpleLogger) Log(message string) {
	fmt.Println(message)
}

// Клиент принимает URL сервера и путь ресурса в качестве аргументов командной строки.
// Сервер принимает адрес и порт в качестве аргументов командной строки.
func parseFlags() (*Config, string, error) {
	cfg := &Config{}
	var requestBody string
	flag.StringVar(&cfg.mode, "mode", "server", "Mode: server or client")
	flag.StringVar(&cfg.address, "address", defaultAddress, "Server address")
	flag.IntVar(&cfg.port, "port", 8080, "Server port")
	flag.StringVar(&cfg.httpMethod, "method", "get", "HTTP method: get or post")
	flag.StringVar(&cfg.resourcePath, "path", "/", "Resource path")
	flag.StringVar(&requestBody, "body", "", "Request body for POST method")

	err := flag.CommandLine.Parse(os.Args[1:])
	if err != nil {
		return nil, "", err
	}

	return cfg, requestBody, nil
}

func main() {
	cfg, requestBody, err := parseFlags()
	if err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	logger := SimpleLogger{}

	switch cfg.mode {
	case "server":
		server := NewServer(cfg.address, cfg.port, logger)
		if err = server.Start(); err != nil {
			logger.Log(fmt.Sprintf("Failed to start server: %v", err))
			os.Exit(1)
		}
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		waitForShutdown(server, sigs)
	case "client":
		client := NewClient(cfg.address, cfg.port, cfg.httpMethod, cfg.resourcePath, logger)

		if cfg.httpMethod == "post" && requestBody != "" {
			var body interface{}
			err = json.Unmarshal([]byte(requestBody), &body)
			if err != nil {
				logger.Log(fmt.Sprintf("Error unmarshalling request body: %v", err))
				os.Exit(1)
			}
			client.Start(body)
		} else {
			client.Start(nil)
		}
	default:
		logger.Log("Unknown mode")
	}
}

type Stoppable interface {
	Stop()
}

func waitForShutdown(server Stoppable, sigs chan os.Signal) {
	<-sigs
	server.Stop()
}
