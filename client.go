package runscope

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/parnurzeal/gorequest"
)

// BaseURL is the Runscope API URL
const BaseURL = "https://api.runscope.com"

// Options used when creating a new client
type Options struct {
	BaseURL string
	Token   string
}

// Client is used when making requests to Runscope
type Client struct {
	req     *gorequest.SuperAgent
	token   string
	baseURL string
}

// Response represents the general response structure returned by Runscope
type Response struct {
	Data  interface{} `json:"data"`
	Error Error       `json:"error"`
	Meta  Meta        `json:"meta"`
}

// Error represents the general error structure returned by Runscope
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Meta represents the general metadata structure returned by Runscope
type Meta struct {
	Status string `json:"status"`
}

func newHTTPClient() *gorequest.SuperAgent {
	client := gorequest.New()
	client.Client = &http.Client{Jar: nil}
	client.Transport = &http.Transport{
		DisableKeepAlives: true,
	}
	return client
}

// NewClient creates a new client for interacting with the Runscope API
func NewClient(options Options) *Client {
	req := newHTTPClient()
	if options.BaseURL == "" {
		options.BaseURL = BaseURL
	}
	return &Client{
		req:     req,
		token:   options.Token,
		baseURL: options.BaseURL,
	}
}

// Get performs a HTTP GET request against the Runscope API
func (client *Client) Get(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)

	client.req.TargetType = "json"
	resp, body, errs := client.req.Get(url).
		Set("Authorization", "Bearer "+client.token).
		EndBytes()
	if errs != nil && len(errs) > 0 {
		return body, errs[len(errs)-1]
	}

	err := checkStatusCode(resp.StatusCode)
	return body, err
}

// Post performs a HTTP POST request against the Rusncope API
// with a supplied payload
func (client *Client) Post(path string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)

	client.req.TargetType = "json"
	resp, body, errs := client.req.Post(url).
		Set("Authorization", "Bearer "+client.token).
		Send(string(data)).
		EndBytes()
	if errs != nil && len(errs) > 0 {
		return body, errs[len(errs)-1]
	}

	err := checkStatusCode(resp.StatusCode)
	return body, err
}

// Put performs a HTTP PUT request against the Runscope API
// with a supplied payload
func (client *Client) Put(path string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)

	client.req.TargetType = "json"
	resp, body, errs := client.req.Put(url).
		Set("Authorization", "Bearer "+client.token).
		Send(string(data)).
		EndBytes()
	if errs != nil && len(errs) > 0 {
		return body, errs[len(errs)-1]
	}

	err := checkStatusCode(resp.StatusCode)
	return body, err
}

// Delete performs a HTTP DELETE request against the Runscope API
func (client *Client) Delete(path string) error {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)

	client.req.TargetType = "json"
	resp, _, errs := client.req.Delete(url).
		Set("Authorization", "Bearer "+client.token).
		EndBytes()
	if errs != nil && len(errs) > 0 {
		return errs[len(errs)-1]
	}

	err := checkStatusCode(resp.StatusCode)
	return err
}

// checkStatusCode returns an error if a HTTP status code does not match 2xx
func checkStatusCode(code int) error {
	ok, err := regexp.MatchString("^2[0-9]{2}$", fmt.Sprintf("%d", code))
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("Request did not match 2xx: %d", code)
	}
	return nil
}
