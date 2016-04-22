package runscope

import (
	"errors"
	"net/http"
)

type Test struct {
	Name                 string        `json:"name"`
	ID                   string        `json:"id"`
	Description          string        `json:"description"`
	CreatedBy            Person        `json:"created_by"`
	CreatedAt            int           `json:"created_at"`
	DefaultEnvironmentID string        `json:"default_environment_id"`
	TriggerURL           string        `json:"trigger_url"`
	LastRun              int           `json:"last_run"`
	Steps                []Step        `json:"steps"`
	Environments         []Environment `json:"environments"`
	Schedules            []Schedule    `json:"schedules"`
}

type UpdateTestRequest struct {
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	DefaultEnvironmentID string   `json:"default_environment_id,omitempty"`
	Steps                []string `json:"steps,omitempty"`
}

func (client *Client) ListTests(bucketKey string) ([]Test, *http.Response, error) {
	var tests = []Test{}
	resp, _, err := client.Get("buckets/"+bucketKey+"/tests", &tests)
	if err != nil {
		println(err.Error())
	}
	return tests, resp, err
}

func (client *Client) GetTest(bucketKey string, testID string) (*Test, *http.Response, error) {
	var test = Test{}
	resp, _, err := client.Get("buckets/"+bucketKey+"/tests/"+testID, &test)
	if err != nil {
		println(err.Error())
	}
	return &test, resp, err
}

func (client *Client) NewTest(bucketKey string, test *Test) (*Test, *http.Response, error) {
	var newTest = Test{}
	if test.Name == "" {
		err := errors.New("Name must not be empty when creating new tests")
		return &newTest, &http.Response{}, err
	}
	resp, _, err := client.Post("buckets/"+bucketKey+"/tests", &test, &newTest)
	if err != nil {
		println(err.Error())
	}
	return &newTest, resp, err
}

func (client *Client) UpdateTest(bucketKey string, testID string, update *UpdateTestRequest) (*Test, *http.Response, error) {
	var newTest = Test{}
	resp, _, err := client.Put("buckets/"+bucketKey+"/tests/"+testID, &update, &newTest)
	if err != nil {
		println(err.Error())
	}
	return &newTest, resp, err
}

func (client *Client) DeleteTest(bucketKey string, testID string) (*http.Response, error) {
	resp, err := client.Delete("buckets/" + bucketKey + "/tests/" + testID)
	if err != nil {
		println(err.Error())
	}
	return resp, err
}
