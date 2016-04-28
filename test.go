package runscope

import (
	"encoding/json"
	"fmt"
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

type NewTestRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTestRequest struct {
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	DefaultEnvironmentID string   `json:"default_environment_id,omitempty"`
	Steps                []string `json:"steps,omitempty"`
}

func (client *Client) ListTests(bucketKey string) (*[]Test, error) {
	var tests = []Test{}

	path := fmt.Sprintf("buckets/%s/tests", bucketKey)

	content, err := client.Get(path)
	if err != nil {
		return &tests, err
	}

	err = unmarshal(content, &tests)
	return &tests, err
}

func (client *Client) GetTest(bucketKey string, testID string) (*Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)

	content, err := client.Get(path)
	if err != nil {
		return &test, err
	}

	err = unmarshal(content, &test)
	return &test, err
}

func (client *Client) NewTest(bucketKey string, newTestRequest *NewTestRequest) (*Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests", bucketKey)
	data, err := json.Marshal(newTestRequest)
	if err != nil {
		return &test, err
	}

	content, err := client.Post(path, data)
	if err != nil {
		return &test, err
	}

	err = unmarshal(content, &test)
	return &test, err
}

func (client *Client) ImportTest(bucketKey string, data []byte) (*Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests", bucketKey)

	content, err := client.Post(path, data)
	if err != nil {
		return &test, err
	}

	err = unmarshal(content, &test)
	return &test, err
}

func (client *Client) UpdateTest(bucketKey string, testID string, update *UpdateTestRequest) (*Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	data, err := json.Marshal(update)
	if err != nil {
		return &test, err
	}

	content, err := client.Put(path, data)
	if err != nil {
		return &test, err
	}

	err = unmarshal(content, &test)
	return &test, err
}

func (client *Client) DeleteTest(bucketKey string, testID string) error {
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	return client.Delete(path)
}
