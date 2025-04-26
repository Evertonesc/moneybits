package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	AuthorizationHeader = "Authorization"
)

type Client struct {
	httpCli  *http.Client
	basePath string
}

func NewRestClient(basePath string) *Client {
	return &Client{
		httpCli: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives:     false,
				MaxIdleConns:          30,
				MaxIdleConnsPerHost:   30,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 5 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DialContext: (&net.Dialer{
					Timeout: 15 * time.Second,
				}).DialContext,
			},
		},
		basePath: basePath,
	}
}

func (c *Client) Get(ctx context.Context, url string, headers, queryParams map[string]string, response any) error {
	req, err := c.buildRequest(ctx, http.MethodGet, c.basePath+url, headers, queryParams, nil)
	if err != nil {
		return err
	}

	return c.do(ctx, req, response)
}

func (c *Client) do(ctx context.Context, request *http.Request, response any) error {
	resp, err := c.httpCli.Do(request)
	if err != nil {
		return err
	}
	defer c.exhaustAndClose(ctx, resp.Body)

	// TODO: handle better with custom errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("http request failed with status code %d: %s", resp.StatusCode, string(errBody))
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (c *Client) buildRequest(ctx context.Context, method, url string, headers, queryParams map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	if body != nil && headers["content-type"] == "" {
		headers["content-type"] = "application/json"
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if len(queryParams) > 0 {
		q := req.URL.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	return req, nil
}

// exhaustAndClose fully reads and discards all remaining bytes from the reader before closing it.
// This prevents TCP connection leaks that can occur when requests are canceled
// before their response bodies are completely read.
func (c *Client) exhaustAndClose(ctx context.Context, r io.ReadCloser) {
	_, err := io.Copy(io.Discard, r)
	if err != nil {
		// handle errors and logs
	}

	err = r.Close()
	if err == nil {
		return
	}

	// in case the body is already closed
	if errors.Is(err, os.ErrClosed) {
		return
	}
}
