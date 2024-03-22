package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run client/main.go <server_url> <resource_path> <http_method> [request_body]")
		os.Exit(1)
	}

	serverURL := os.Args[1]
	resourcePath := os.Args[2]
	httpMethod := os.Args[3]
	var requestBody []byte
	if len(os.Args) > 4 {
		requestBody = []byte(os.Args[4])
	}

	client := &http.Client{}
	var req *http.Request
	var err error

	switch httpMethod {
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, serverURL+resourcePath, nil)
	case http.MethodPost:
		req, err = http.NewRequest(http.MethodPost, serverURL+resourcePath, bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
	case http.MethodPut:
		req, err = http.NewRequest(http.MethodPut, serverURL+resourcePath, bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
	case http.MethodDelete:
		req, err = http.NewRequest(http.MethodDelete, serverURL+resourcePath+"?id="+os.Args[4], nil)
	default:
		fmt.Printf("Unsupported HTTP method: %s\n", httpMethod)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
