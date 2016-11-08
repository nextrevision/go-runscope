package runscope

import (
	"encoding/json"
	"fmt"
)

// Environment represents a Runscope Environment
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

// Email represents an email notification setting for a test or event
type Email struct {
	NotifyAll       bool     `json:"notify_all"`
	NotifyOn        string   `json:"notify_on"`
	NotifyThreshold int      `json:"notify_threshold"`
	Recipients      []Person `json:"recipients"`
}

// RemoteAgent represents Runscope remote agents
type RemoteAgent struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// ListTestEnvironments returns all environments associated with a given test
func (client *Client) ListTestEnvironments(bucketKey string, testID string) ([]Environment, error) {
	var environments = []Environment{}

	path := fmt.Sprintf("buckets/%s/tests/%s/environments", bucketKey, testID)
	content, err := client.Get(path)
	if err != nil {
		return environments, err
	}

	err = unmarshal(content, &environments)
	return environments, err
}

// ListSharedEnvironments returns shared environments in a given bucket
func (client *Client) ListSharedEnvironments(bucketKey string) ([]Environment, error) {
	var environments = []Environment{}

	path := fmt.Sprintf("buckets/%s/environments", bucketKey)
	content, err := client.Get(path)
	if err != nil {
		return environments, err
	}

	err = unmarshal(content, &environments)
	return environments, err
}

// GetTestEnvironment fetches the details for a given
// environment associated with a test
func (client *Client) GetTestEnvironment(bucketKey string, testID string, environmentID string) (Environment, error) {
	var environment = Environment{}

	path := fmt.Sprintf("buckets/%s/tests/%s/environments/%s", bucketKey, testID, environmentID)
	content, err := client.Get(path)
	if err != nil {
		return environment, err
	}

	err = unmarshal(content, &environment)
	return environment, err
}

// GetSharedEnvironment fetches the details of a given
// environment assocaited with a bucket
func (client *Client) GetSharedEnvironment(bucketKey string, environmentID string) (Environment, error) {
	var environment = Environment{}

	path := fmt.Sprintf("buckets/%s/environments/%s", bucketKey, environmentID)
	content, err := client.Get(path)
	if err != nil {
		return environment, err
	}

	err = unmarshal(content, &environment)
	return environment, err
}

// NewTestEnvironment creates an environment for a given test
func (client *Client) NewTestEnvironment(bucketKey string, testID string, environment Environment) (Environment, error) {
	var newEnvironment = Environment{}

	path := fmt.Sprintf("buckets/%s/tests/%s/environments", bucketKey, testID)
	data, err := json.Marshal(&environment)
	if err != nil {
		return newEnvironment, err
	}

	content, err := client.Post(path, data)
	if err != nil {
		return newEnvironment, err
	}

	err = unmarshal(content, &newEnvironment)
	return newEnvironment, err
}

// NewSharedEnvironment creates a new shared environment in a bucket
func (client *Client) NewSharedEnvironment(bucketKey string, environment Environment) (Environment, error) {
	var newEnvironment = Environment{}

	path := fmt.Sprintf("buckets/%s/environments", bucketKey)
	data, err := json.Marshal(&environment)
	if err != nil {
		return newEnvironment, err
	}

	content, err := client.Post(path, data)
	if err != nil {
		return newEnvironment, err
	}

	err = unmarshal(content, &newEnvironment)
	return newEnvironment, err
}

// UpdateTestEnvironment updates a test environment
func (client *Client) UpdateTestEnvironment(bucketKey string, testID string, environmentID string, environment Environment) (Environment, error) {
	var newEnvironment = Environment{}

	path := fmt.Sprintf("buckets/%s/tests/%s/environments/%s", bucketKey, testID, environmentID)
	data, err := json.Marshal(&environment)
	if err != nil {
		return newEnvironment, err
	}

	content, err := client.Put(path, data)
	if err != nil {
		return newEnvironment, err
	}

	err = unmarshal(content, &newEnvironment)
	return newEnvironment, err
}

// UpdateSharedEnvironment updates a shared environment in a bucket
func (client *Client) UpdateSharedEnvironment(bucketKey string, environmentID string, environment Environment) (Environment, error) {
	var newEnvironment = Environment{}

	path := fmt.Sprintf("buckets/%s/environments/%s", bucketKey, environmentID)
	data, err := json.Marshal(&environment)
	if err != nil {
		return newEnvironment, err
	}

	content, err := client.Put(path, data)
	if err != nil {
		return newEnvironment, err
	}

	err = unmarshal(content, &newEnvironment)
	return newEnvironment, err
}

// DeleteEnvironment removes an environment from a bucket
func (client *Client) DeleteEnvironment(bucketKey string, environmentID string) error {
	path := fmt.Sprintf("buckets/%s/environments/%s", bucketKey, environmentID)
	return client.Delete(path)
}
