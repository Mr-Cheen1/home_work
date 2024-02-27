package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	address      string
	port         int
	httpMethod   string
	resourcePath string
	logger       Logger
}

func NewClient(address string, port int, httpMethod, resourcePath string, logger Logger) Client {
	return Client{address, port, httpMethod, resourcePath, logger}
}

func (c Client) Start() {
	url := fmt.Sprintf("http://%s:%d%s", c.address, c.port, c.resourcePath)
	client := &http.Client{}
	var resp *http.Response
	var req *http.Request
	var err error

	method := strings.ToLower(c.httpMethod)
	// Клиент отправляет HTTP GET и POST запросы.
	switch method {
	case "get":
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			c.logger.Log(fmt.Sprintf("Error creating GET request for %s: %v", url, err))
			return
		}
	case "post":
		req, err = http.NewRequest("POST", url, nil)
		if err != nil {
			c.logger.Log(fmt.Sprintf("Error creating POST request for %s: %v", url, err))
			return
		}
		req.Header.Set("Content-Type", "application/json")
	default:
		errorMessage := fmt.Sprintf("Unsupported method: %s, Status code: %d", c.httpMethod, http.StatusMethodNotAllowed)
		c.logger.Log(errorMessage)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		c.logger.Log(fmt.Sprintf("Client error on %s request to %s: %v", strings.ToUpper(c.httpMethod), url, err))
		return
	}
	defer resp.Body.Close()

	// Клиент выводит полученный ответ от сервера в стандартный поток вывода.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Log(fmt.Sprintf("Error reading response from %s: %v", url, err))
		return
	}

	c.logger.Log(fmt.Sprintf("Client received response from %s: %s, Status code: %d", url, string(body), resp.StatusCode))
}
