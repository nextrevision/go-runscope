package runscope

import (
	"fmt"
	"net/http"
	"regexp"

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
