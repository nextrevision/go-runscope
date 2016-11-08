package runscope

import (
	"encoding/json"
	"fmt"
)

// Test represents a Runscope test
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

// LastRun represents the last result of a test
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

// NewTestRequest represents all parameters for creating a new test
type NewTestRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateTestRequest represents all parameters for updating an existing test
type UpdateTestRequest struct {
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	DefaultEnvironmentID string   `json:"default_environment_id,omitempty"`
	Steps                []string `json:"steps,omitempty"`
}

// ListTestOptions are parameters for modifying the tests returned
type ListTestOptions struct {
	Count  int
	Offset int
}

// ListTests returns a slice of Tests for a given bucket
func (client *Client) ListTests(bucketKey string, options ListTestOptions) ([]Test, error) {
	var tests = []Test{}

	path := fmt.Sprintf("buckets/%s/tests", bucketKey)

	if options.Count != 0 {
		path = fmt.Sprintf("%s?count=%d&offset=%d", path, options.Count, options.Offset)
	}

	content, err := client.Get(path)
	if err != nil {
		return tests, err
	}

	err = unmarshal(content, &tests)
	return tests, err
}

// GetTest returns details about a given test
func (client *Client) GetTest(bucketKey string, testID string) (Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)

	content, err := client.Get(path)
	if err != nil {
		return test, err
	}

	err = unmarshal(content, &test)
	return test, err
}

// NewTest creates a new test in a given bucket
func (client *Client) NewTest(bucketKey string, newTestRequest NewTestRequest) (Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests", bucketKey)
	data, err := json.Marshal(&newTestRequest)
	if err != nil {
		return test, err
	}

	content, err := client.Post(path, data)
	if err != nil {
		return test, err
	}

	err = unmarshal(content, &test)
	return test, err
}

// UpdateTest modifies an existing test in a given bucket
func (client *Client) UpdateTest(bucketKey string, testID string, updateTestRequest UpdateTestRequest) (Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	data, err := json.Marshal(updateTestRequest)
	if err != nil {
		return test, err
	}

	content, err := client.Put(path, data)
	if err != nil {
		return test, err
	}

	err = unmarshal(content, &test)
	return test, err
}

// ImportTest creates a test for a given bucket with a JSON payload
func (client *Client) ImportTest(bucketKey string, data []byte) (Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests", bucketKey)

	content, err := client.Post(path, data)
	if err != nil {
		return test, err
	}

	err = unmarshal(content, &test)
	return test, err
}

// ReimportTest updates an existing test for a given bucket with a JSON payload
func (client *Client) ReimportTest(bucketKey string, testID string, data []byte) (Test, error) {
	var test = Test{}

	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)

	content, err := client.Put(path, data)
	if err != nil {
		return test, err
	}

	err = unmarshal(content, &test)
	return test, err
}

// DeleteTest removes a test from a bucket
func (client *Client) DeleteTest(bucketKey string, testID string) error {
	path := fmt.Sprintf("buckets/%s/tests/%s", bucketKey, testID)
	return client.Delete(path)
}
