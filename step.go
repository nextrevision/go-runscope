package runscope

import (
	"encoding/json"
	"fmt"
)

// Step represents a single test step
type Step struct {
	StepType         string              `json:"step_type"`
	ID               string              `json:"id"`
	Method           string              `json:"method"`
	URL              string              `json:"url"`
	Body             string              `json:"body"`
	Auth             Auth                `json:"auth"`
	Form             map[string][]string `json:"form"`
	Assertions       []Assertion         `json:"assertions"`
	Variables        []Variable          `json:"variables"`
	Headers          map[string][]string `json:"headers"`
	Scripts          []string            `json:"scripts"`
	Note             string              `json:"note"`
	Duration         int                 `json:"duration"`
	Comparison       string              `json:"comparison"`
	RightValue       string              `json:"right_value"`
	LeftValue        string              `json:"left_value"`
	Steps            []Step              `json:"steps"`
	IntegrationID    string              `json:"integration_id"`
	SuiteID          string              `json:"suite_id"`
	TestID           string              `json:"test_id"`
	IsCustomStartURL bool                `json:"is_custom_start_url"`
}

// Assertion represents a Runscope test assertion
type Assertion struct {
	Source      string      `json:"source"`
	Property    string      `json:"property"`
	Comparison  string      `json:"comparison"`
	Value       interface{} `json:"value"`
	Result      string      `json:"result,omitempty"`
	TargetValue interface{} `json:"target_value,omitempty"`
	ActualValue interface{} `json:"actual_value,omitempty"`
	Error       string      `json:"error,omitempty"`
}

// Auth represents auth parameters for a test
type Auth struct {
	AuthType       string `json:"auth_type"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"username,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
	TokenSecret    string `json:"token_secret,omitempty"`
	ConsumerKey    string `json:"consumer_key,omitempty"`
	ConsumerSecret string `json:"consumer_secret,omitempty"`
	SignatureType  string `json:"signature_type,omitempty"`
}

// Variable represents a Runscope variable
type Variable struct {
	Name     string      `json:"name"`
	Source   string      `json:"source"`
	Property string      `json:"property"`
	Result   string      `json:"result,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Error    string      `json:"error,omitempty"`
}

// Script represents a Runscope script resource
// Test details will list scripts as []string's instead of
// the format of this struct
type Script struct {
	Value  string `json:"value"`
	Result string `json:"result,omitempty"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

// ListSteps returns all steps for a given test
func (client *Client) ListSteps(bucketKey string, testID string) ([]Step, error) {
	var steps = []Step{}

	path := fmt.Sprintf("buckets/%s/tests/%s/steps", bucketKey, testID)

	content, err := client.Get(path)
	if err != nil {
		return steps, err
	}

	err = unmarshal(content, &steps)
	return steps, err
}

// GetStep returns details for a given step
func (client *Client) GetStep(bucketKey string, testID string, stepID string) (Step, error) {
	var step = Step{}

	path := fmt.Sprintf("buckets/%s/tests/%s/steps/%s", bucketKey, testID, stepID)
	content, err := client.Get(path)
	if err != nil {
		return step, err
	}

	err = unmarshal(content, &step)
	return step, err
}

// NewStep creates a new step for a given test
func (client *Client) NewStep(bucketKey string, testID string, step Step) (Step, error) {
	var newStep = Step{}

	path := fmt.Sprintf("buckets/%s/tests/%s/steps", bucketKey, testID)
	data, err := json.Marshal(&step)
	if err != nil {
		return newStep, err
	}

	content, err := client.Post(path, data)
	if err != nil {
		return newStep, err
	}

	err = unmarshal(content, &newStep)
	return newStep, err
}

// UpdateStep updates an existing step according to the Step struct passed to it
func (client *Client) UpdateStep(bucketKey string, testID string, stepID string, step Step) (Step, error) {
	var newStep = Step{}

	path := fmt.Sprintf("buckets/%s/tests/%s/steps/%s", bucketKey, testID, stepID)
	data, err := json.Marshal(&step)
	if err != nil {
		return newStep, err
	}

	content, err := client.Put(path, data)
	if err != nil {
		return newStep, err
	}

	err = unmarshal(content, &newStep)
	return newStep, err
}

// DeleteStep removes a step from a test
func (client *Client) DeleteStep(bucketKey string, testID string, stepID string) error {
	path := fmt.Sprintf("buckets/%s/tests/%s/steps/%s", bucketKey, testID, stepID)
	return client.Delete(path)
}
