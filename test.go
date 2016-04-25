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
	LastRun              LastRun       `json:"last_run"`
	Steps                []Step        `json:"steps"`
	Environments         []Environment `json:"environments"`
	Schedules            []Schedule    `json:"schedules"`
}

type LastRun struct {
	ID                 string   `json:"id"`
	UUID               string   `json:"uuid"`
	TestUUID           string   `json:"test_uuid"`
	EnvironmentUUID    string   `json:"environment_uuid"`
	EnvironmentName    string   `json:"environment_name"`
	RemoteAgentUUID    string   `json:"remote_agent_uuid"`
	RemoteAgentName    string   `json:"remote_agent_name"`
	RemoteAgentVersion string   `json:"remote_agent_version"`
	Status             string   `json:"status"`
	CreatedAt          float64  `json:"created_at"`
	FinishedAt         float64  `json:"finished_at"`
	ErrorCount         int      `json:"error_count"`
	MessageSuccess     int      `json:"message_success"`
	Source             string   `json:"source"`
	ExtractorCount     int      `json:"extractor_count"`
	ExtractorSuccess   int      `json:"extractor_success"`
	SubstituionCount   int      `json:"substitution_count"`
	SubstituionSuccess int      `json:"substitution_success"`
	ScriptCount        int      `json:"script_count"`
	ScriptSuccess      int      `json:"script_success"`
	AssertionCount     int      `json:"assertion_count"`
	AssertionSuccess   int      `json:"assertion_success"`
	BucketKey          string   `json:"bucket_key"`
	Region             string   `json:"region"`
	Messages           []string `json:"messages"`
	MessageCount       int      `json:"message_count"`
	TemplateUUIDs      []string `json:"template_uuids"`
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
	resp, err := client.Get(path, &tests)
	return &tests, resp, err
}

func (client *Client) GetTest(bucketKey string, testID string) (*Test, *http.Response, error) {
	var test = Test{}
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	resp, err := client.Get(path, &test)
	return &test, resp, err
}

func (client *Client) NewTest(bucketKey string, test *Test) (*Test, *http.Response, error) {
	var newTest = Test{}
	if test.Name == "" {
		err := errors.New("Name must not be empty when creating new tests")
		return &newTest, &http.Response{}, err
	}
	path := fmt.Sprintf("buckets/%s/tests", bucketKey)
	resp, err := client.Post(path, &test, &newTest)
	return &newTest, resp, err
}

func (client *Client) UpdateTest(bucketKey string, testID string, update *UpdateTestRequest) (*Test, *http.Response, error) {
	var newTest = Test{}
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	resp, err := client.Put(path, &update, &newTest)
	return &newTest, resp, err
}

func (client *Client) DeleteTest(bucketKey string, testID string) (*http.Response, error) {
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	resp, err := client.Delete(path)
	return resp, err
}
