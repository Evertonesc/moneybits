//go:build integration

package setup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestClient struct {
	Server *httptest.Server
}

func NewTestRestClient(server *http.Server) *TestClient {
	if server == nil {
		panic("server cannot be nil")
	}

	handler := server.Handler
	if handler == nil {
		handler = http.DefaultServeMux
	}

	return &TestClient{
		Server: httptest.NewServer(handler),
	}
}

func (c *TestClient) Post(path string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	resp, err := http.Post(c.Server.URL+path, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	return resp, nil
}

func (c *TestClient) Get(path string) (*http.Response, error) {
	resp, err := http.Get(c.Server.URL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	return resp, nil
}

func (c *TestClient) Put(path string, body interface{}, response interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, c.Server.URL+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	if response != nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}

		if err := json.Unmarshal(body, response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %v", err)
		}
	}

	return resp, nil
}
