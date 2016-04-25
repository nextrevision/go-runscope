package runscope

import (
	"fmt"
	"net/http"
)

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
	Scripts          []Script            `json:"scripts"`
	Note             string              `json:"note"`
	Duration         int                 `json:"duration"`
	Comparision      string              `json:"string"`
	RightValue       string              `json:"right_value"`
	LeftValue        string              `json:"left_value"`
	Steps            []Step              `json:"steps"`
	IntegrationID    string              `json:"integration_id"`
	SuiteID          string              `json:"suite_id"`
	TestID           string              `json:"test_id"`
	IsCustomStartURL bool                `json:"is_custom_start_url"`
}

type Assertion struct {
	Source     string      `json:"source"`
	Property   string      `json:"property"`
	Comparison string      `json:"comparison"`
	Value      interface{} `json:"value"`
}

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

type Variable struct {
	Name     string `json:"name"`
	Source   string `json:"source"`
	Property string `json:"property"`
}

type Script struct {
	Value string `json:"value"`
}

func (client *Client) ListSteps(bucketKey string, testID string) (*[]Step, *http.Response, error) {
	var steps = []Step{}
	path := fmt.Sprintf("buckets/%s/tests/%s/steps", bucketKey, testID)
	resp, err := client.Get(path, &steps)
	return &steps, resp, err
}

func (client *Client) GetStep(bucketKey string, testID string, stepID string) (*Step, *http.Response, error) {
	var step = Step{}
	path := fmt.Sprintf("buckets/%s/tests/%s/steps/%s", bucketKey, testID, stepID)
	resp, err := client.Get(path, &step)
	return &step, resp, err
}

func (client *Client) NewStep(bucketKey string, testID string, step *Step) (*Step, *http.Response, error) {
	var newStep = Step{}
	path := fmt.Sprintf("buckets/%s/tests/%s/steps", bucketKey, testID)
	resp, err := client.Post(path, &step, &newStep)
	return &newStep, resp, err
}

func (client *Client) UpdateStep(bucketKey string, testID string, stepID string, step *Step) (*Step, *http.Response, error) {
	var newStep = Step{}
	path := fmt.Sprintf("buckets/%s/tests/%s/steps/%s", bucketKey, testID, stepID)
	resp, err := client.Put(path, &step, &newStep)
	return &newStep, resp, err
}

func (client *Client) DeleteStep(bucketKey string, testID string, stepID string) (*http.Response, error) {
	path := fmt.Sprintf("buckets/%s/tests/%s/steps/%s", bucketKey, testID, stepID)
	resp, err := client.Delete(path)
	return resp, err
}
