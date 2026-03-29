package polydata

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	json "github.com/go-json-experiment/json"
)

const DefaultHost = "https://data-api.polymarket.com"

type Client struct {
	host      string
	http      *http.Client
	userAgent string
}

type Config struct {
	Host       string
	HTTPClient *http.Client
	UserAgent  string
}

func New(config Config) *Client {
	if config.Host == "" {
		config.Host = DefaultHost
	}
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	if config.UserAgent == "" {
		config.UserAgent = "go-polymarket-data"
	}
	return &Client{
		host:      config.Host,
		http:      config.HTTPClient,
		userAgent: config.UserAgent,
	}
}

func (c *Client) getJSON(ctx context.Context, path string, query url.Values, out any) error {
	fullURL := c.host + path
	if len(query) > 0 {
		fullURL += "?" + query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		payload, _ := io.ReadAll(resp.Body)
		return &APIError{
			StatusCode: resp.StatusCode,
			Body:       payload,
			Message:    strings.TrimSpace(string(payload)),
		}
	}

	if out == nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil
	}

	if err := json.UnmarshalRead(resp.Body, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

type APIError struct {
	StatusCode int
	Message    string
	Body       []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("polymarket data API error: status %d: %s", e.StatusCode, e.Message)
}

func (e *APIError) HTTPStatus() int { return e.StatusCode }

func setBool(query url.Values, key string, val *bool) {
	if val != nil {
		query.Set(key, strconv.FormatBool(*val))
	}
}

func setInt(query url.Values, key string, val int) {
	if val > 0 {
		query.Set(key, strconv.Itoa(val))
	}
}

func setInt64(query url.Values, key string, val int64) {
	if val > 0 {
		query.Set(key, strconv.FormatInt(val, 10))
	}
}

func setString(query url.Values, key, val string) {
	if val != "" {
		query.Set(key, val)
	}
}

func setCommaList(query url.Values, key string, vals []string) {
	if len(vals) > 0 {
		query.Set(key, strings.Join(vals, ","))
	}
}
