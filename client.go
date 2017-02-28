package runscope

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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
	*http.Client
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

func newHTTPClient() *http.Client {
	return &http.Client{
		Jar:       nil,
		Transport: &http.Transport{DisableKeepAlives: true},
	}
}

// NewClient creates a new client for interacting with the Runscope API
func NewClient(options Options) *Client {
	client := newHTTPClient()
	if options.BaseURL == "" {
		options.BaseURL = BaseURL
	}
	return &Client{
		Client:  client,
		token:   options.Token,
		baseURL: options.BaseURL,
	}
}

// Get performs a HTTP GET request against the Runscope API
func (client *Client) Get(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	return client.doRequest("GET", url, nil)
}

// Post performs a HTTP POST request against the Rusncope API
// with a supplied payload
func (client *Client) Post(path string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	return client.doRequest("POST", url, data)
}

// Put performs a HTTP PUT request against the Runscope API
// with a supplied payload
func (client *Client) Put(path string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	return client.doRequest("PUT", url, data)
}

// Delete performs a HTTP DELETE request against the Runscope API
func (client *Client) Delete(path string) error {
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	_, err := client.doRequest("DELETE", url, nil)
	return err
}

// checkStatusCode returns an error if a HTTP status code does not match 2xx
func checkStatusCode(code int) error {
	if 200 <= code && code < 300 {
		return nil
	}
	return fmt.Errorf("Request did not match 2xx: %d", code)
}

func (client *Client) doRequest(method string, url string, data []byte) ([]byte, error) {
	reqBody := bytes.NewReader(data)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	setHeaders(req, client.token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	err = checkStatusCode(res.StatusCode)
	return body, err
}

func setHeaders(req *http.Request, token string) {
	req.Header = map[string][]string{
		"Authorization": {"Bearer " + token},
		"Accept":        {"application/json"},
		"Content-Type":  {"application/json"},
	}
}
