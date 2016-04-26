package runscope

import (
	"fmt"
	"net/http"
)

type Environment struct {
	Name                string            `json:"name"`
	ID                  string            `json:"id"`
	ParentEnvironmentID string            `json:"parent_environment_id"`
	PreserveCookies     bool              `json:"preserve_cookies"`
	Regions             []string          `json:"regions"`
	RemoteAgents        []RemoteAgent     `json:"remote_agents"`
	Script              string            `json:"script"`
	TestID              string            `json:"test_id"`
	VerifySSL           bool              `json:"verify_ssl"`
	Webhooks            []string          `json:"webhooks"`
	Emails              Email             `json:"emails"`
	InitialVariables    map[string]string `json:"initial_variables"`
	Integrations        []TeamIntegration `json:"integrations"`
}

type Email struct {
	NotifyAll       bool     `json:"notify_all"`
	NotifyOn        string   `json:"notify_on"`
	NotifyThreshold int      `json:"notify_threshold"`
	Recipients      []Person `json:"recipients"`
}

type RemoteAgent struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

func (client *Client) ListTestEnvironments(bucketKey string, testID string) (*[]Environment, *http.Response, error) {
	var environments = []Environment{}
	path := fmt.Sprintf("buckets/%s/tests/%s/environments", bucketKey, testID)
	resp, err := client.Get(path, &environments)
	return &environments, resp, err
}

func (client *Client) ListSharedEnvironments(bucketKey string) (*[]Environment, *http.Response, error) {
	var environments = []Environment{}
	path := fmt.Sprintf("buckets/%s/environments", bucketKey)
	resp, err := client.Get(path, &environments)
	return &environments, resp, err
}

func (client *Client) GetTestEnvironment(bucketKey string, testID string, environmentID string) (*Environment, *http.Response, error) {
	var environment = Environment{}
	path := fmt.Sprintf("buckets/%s/tests/%s/environments/%s", bucketKey, testID, environmentID)
	resp, err := client.Get(path, &environment)
	return &environment, resp, err
}

func (client *Client) GetSharedEnvironment(bucketKey string, environmentID string) (*Environment, *http.Response, error) {
	var environment = Environment{}
	path := fmt.Sprintf("buckets/%s/environments/%s", bucketKey, environmentID)
	resp, err := client.Get(path, &environment)
	return &environment, resp, err
}

func (client *Client) NewTestEnvironment(bucketKey string, testID string, environment *Environment) (*Environment, *http.Response, error) {
	var newEnvironment = Environment{}
	path := fmt.Sprintf("buckets/%s/tests/%s/environments", bucketKey, testID)
	resp, err := client.Post(path, &environment, &newEnvironment)
	return &newEnvironment, resp, err
}

func (client *Client) NewSharedEnvironment(bucketKey string, environment *Environment) (*Environment, *http.Response, error) {
	var newEnvironment = Environment{}
	path := fmt.Sprintf("buckets/%s/environments", bucketKey)
	resp, err := client.Post(path, &environment, &newEnvironment)
	return &newEnvironment, resp, err
}

func (client *Client) UpdateTestEnvironment(bucketKey string, testID string, environmentID string, environment *Environment) (*Environment, *http.Response, error) {
	var newEnvironment = Environment{}
	path := fmt.Sprintf("buckets/%s/tests/%s/environments/%s", bucketKey, testID, environmentID)
	resp, err := client.Put(path, &environment, &newEnvironment)
	return &newEnvironment, resp, err
}

func (client *Client) UpdateSharedEnvironment(bucketKey string, environmentID string, environment *Environment) (*Environment, *http.Response, error) {
	var newEnvironment = Environment{}
	path := fmt.Sprintf("buckets/%s/environments/%s", bucketKey, environmentID)
	resp, err := client.Put(path, &environment, &newEnvironment)
	return &newEnvironment, resp, err
}

func (client *Client) DeleteEnvironment(bucketKey string, environmentID string) (*http.Response, error) {
	path := fmt.Sprintf("buckets/%s/environments/%s", bucketKey, environmentID)
	resp, err := client.Delete(path)
	return resp, err
}
