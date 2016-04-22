package runscope

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// BaseURL struct
const BaseURL = "https://api.runscope.com"

// Options struct
type Options struct {
	BaseURL string
	Token   string
}

// Client struct
type Client struct {
	req     *gorequest.SuperAgent
	token   string
	baseURL string
}

// Response struct
type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

// Meta struct
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

// NewClient returns a new Runscope Client object
func NewClient(options *Options) *Client {
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

// Get returns a basic Response object
func (client *Client) Get(path string, result interface{}) (*http.Response, *Response, error) {
	var response = Response{Data: result}
	client.req.TargetType = "json"
	resp, body, errs := client.req.Get(fmt.Sprintf("%s/%s", client.baseURL, path)).Set("Authorization", "Bearer "+client.token).EndBytes()
	if errs != nil && len(errs) > 0 {
		return nil, &response, errs[len(errs)-1]
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &response, err
	}
	_resp := http.Response(*resp)
	return &_resp, &response, nil
}

// Post returns a basic Response object
func (client *Client) Post(path string, data interface{}, result interface{}) (*http.Response, *Response, error) {
	var response = Response{Data: result}
	client.req.TargetType = "json"
	resp, body, errs := client.req.Post(fmt.Sprintf("%s/%s", client.baseURL, path)).Set("Authorization", "Bearer "+client.token).SendStruct(data).EndBytes()
	if errs != nil && len(errs) > 0 {
		return nil, &response, errs[len(errs)-1]
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &response, err
	}
	_resp := http.Response(*resp)
	return &_resp, &response, nil
}

// Delete TODO
func (client *Client) Delete(path string) (*http.Response, error) {
	client.req.TargetType = "json"
	resp, _, errs := client.req.Delete(fmt.Sprintf("%s/%s", client.baseURL, path)).Set("Authorization", "Bearer "+client.token).EndBytes()
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	_resp := http.Response(*resp)
	return &_resp, nil
}
