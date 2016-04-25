package runscope

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

const BaseURL = "https://api.runscope.com"

type Options struct {
	BaseURL string
	Token   string
}

type Client struct {
	req     *gorequest.SuperAgent
	token   string
	baseURL string
}

type Response struct {
	Data  interface{} `json:"data"`
	Error Error       `json:"error"`
	Meta  Meta        `json:"meta"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

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

func (client *Client) Get(path string, result interface{}) (*http.Response, error) {
	var response = Response{Data: result}
	client.req.TargetType = "json"
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	resp, body, errs := client.req.Get(url).Set("Authorization", "Bearer "+client.token).EndBytes()
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	_resp := http.Response(*resp)
	return &_resp, nil
}

func (client *Client) Post(path string, data interface{}, result interface{}) (*http.Response, error) {
	var response = Response{Data: result}
	client.req.TargetType = "json"
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	resp, body, errs := client.req.Post(url).Set("Authorization", "Bearer "+client.token).SendStruct(data).EndBytes()
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	_resp := http.Response(*resp)
	return &_resp, nil
}

func (client *Client) Put(path string, data interface{}, result interface{}) (*http.Response, error) {
	var response = Response{Data: result}
	client.req.TargetType = "json"
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	resp, body, errs := client.req.Put(url).Set("Authorization", "Bearer "+client.token).SendStruct(data).EndBytes()
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	_resp := http.Response(*resp)
	return &_resp, nil
}

func (client *Client) Delete(path string) (*http.Response, error) {
	client.req.TargetType = "json"
	url := fmt.Sprintf("%s/%s", client.baseURL, path)
	resp, _, errs := client.req.Delete(url).Set("Authorization", "Bearer "+client.token).EndBytes()
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	_resp := http.Response(*resp)
	return &_resp, nil
}
