package main

import (
	"bytes"
	"encoding/json"
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

// Добавлен параметр body в метод Start, который может быть любым типом данных.
// Этот параметр будет сериализован в JSON и отправлен в теле POST запроса.
func (c *Client) Start(body interface{}) {
	url := fmt.Sprintf("http://%s:%d%s", c.address, c.port, c.resourcePath)
	client := &http.Client{}
	var resp *http.Response
	var req *http.Request
	var err error

	// Сериализация тела запроса в JSON.
	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.logger.Log(fmt.Sprintf("Error marshalling request body: %v", err))
		return
	}

	method := strings.ToLower(c.httpMethod)
	switch method {
	case "get":
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			c.logger.Log(fmt.Sprintf("Error creating GET request for %s: %v", url, err))
			return
		}
	case "post":
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
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

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Log(fmt.Sprintf("Error reading response from %s: %v", url, err))
		return
	}

	c.logger.Log(fmt.Sprintf("Client received response from %s: %s, Status code: %d",
		url, string(responseBody), resp.StatusCode))
}
