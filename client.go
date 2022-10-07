package gohypixel

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/efectn/go-hypixel/models"
)

type Config struct {
	HTTPClient *http.Client
	UserAgent  string
	APIKey     string
}

type Client struct {
	Client    *http.Client
	UserAgent string
	APIKey    string
}

var defaultConfig = Config{
	HTTPClient: &http.Client{Timeout: 5 * time.Second},
	UserAgent:  "",
}

var apiURL = "https://api.hypixel.net/"

// New is a function to return Hypixel client instance.
func New(config Config) *Client {
	// Validate config
	if config.APIKey == "" {
		panic("go-hypixel: you must enter API key!")
	}

	if config.HTTPClient == nil {
		config.HTTPClient = defaultConfig.HTTPClient
	}

	if config.UserAgent == "" {
		config.UserAgent = defaultConfig.UserAgent
	}

	return &Client{
		Client:    config.HTTPClient,
		UserAgent: config.UserAgent,
		APIKey:    config.APIKey,
	}
}

func (c *Client) generateReq(path string) (string, *http.Response, error) {
	// Create new request
	req, err := http.NewRequest("GET", apiURL+path, nil)
	if err != nil {
		return "", nil, err
	}

	// Set agent & necessary headers
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("API-Key", c.APIKey)

	// Send request
	resp, err := c.Client.Do(req)
	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	// Return body, resp
	return string(body), resp, nil
}

func (c *Client) Key() (models.Key, error) {
	var key models.Key

	// Send request
	body, resp, err := c.generateReq("key")
	if err != nil {
		return key, err
	}

	// Check HTTP errors
	if err := checkHTTPErrors(resp, body); err != nil {
		return key, err
	}

	if err := json.Unmarshal([]byte(body), &key); err != nil {
		return key, err
	}

	return key, nil
}
