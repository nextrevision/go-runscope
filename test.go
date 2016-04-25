package runscope

import (
	"errors"
	"fmt"
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

func (client *Client) ListTests(bucketKey string) (*[]Test, *http.Response, error) {
	var tests = []Test{}
	path := fmt.Sprintf("buckets/%s/tests", bucketKey)
	resp, _, err := client.Get(path, &tests)
	if err != nil {
		println(err.Error())
	}
	return &tests, resp, err
}

func (client *Client) GetTest(bucketKey string, testID string) (*Test, *http.Response, error) {
	var test = Test{}
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	resp, _, err := client.Get(path, &test)
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
	path := fmt.Sprintf("buckets/%s/tests", bucketKey)
	resp, _, err := client.Post(path, &test, &newTest)
	if err != nil {
		println(err.Error())
	}
	return &newTest, resp, err
}

func (client *Client) UpdateTest(bucketKey string, testID string, update *UpdateTestRequest) (*Test, *http.Response, error) {
	var newTest = Test{}
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	resp, _, err := client.Put(path, &update, &newTest)
	if err != nil {
		println(err.Error())
	}
	return &newTest, resp, err
}

func (client *Client) DeleteTest(bucketKey string, testID string) (*http.Response, error) {
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	resp, err := client.Delete(path)
	if err != nil {
		println(err.Error())
	}
	return resp, err
}
