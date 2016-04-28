package runscope

import "fmt"

type Result struct {
	AssertionsDefined int       `json:"assertions_defined"`
	AssertionsFailed  int       `json:"assertions_failed"`
	AssertionsPassed  int       `json:"assertions_passed"`
	BucketKey         string    `json:"bucket_key"`
	FinishedAt        float64   `json:"finished_at"`
	Region            string    `json:"region"`
	RequestsExecuted  int       `json:"requests_executed"`
	Result            string    `json:"result"`
	ScriptsDefined    int       `json:"scripts_defined"`
	ScriptsFailed     int       `json:"scripts_failed"`
	ScriptsPassed     int       `json:"scripts_passed"`
	StartedAt         float64   `json:"started_at"`
	TestRunID         string    `json:"test_run_id"`
	TestRunURL        string    `json:"test_run_url"`
	TestID            string    `json:"test_id"`
	VariablesDefined  int       `json:"variables_defined"`
	VariablesFailed   int       `json:"variables_failed"`
	VariablesPassed   int       `json:"variables_passed"`
	EnvironmentID     string    `json:"environment_id"`
	EnvironmentName   string    `json:"environment_name"`
	Requests          []Request `json:"requests"`
}

type Request struct {
	Result            string      `json:"result"`
	URL               string      `json:"url"`
	Method            string      `json:"method"`
	AssertionsDefined int         `json:"assertions_defined"`
	AssertionsFailed  int         `json:"assertions_failed"`
	AssertionsPassed  int         `json:"assertions_passed"`
	ScriptsDefined    int         `json:"scripts_defined"`
	ScriptsFailed     int         `json:"scripts_failed"`
	ScriptsPassed     int         `json:"scripts_passed"`
	VariablesDefined  int         `json:"variables_defined"`
	VariablesFailed   int         `json:"variables_failed"`
	VariablesPassed   int         `json:"variables_passed"`
	Assertions        []Assertion `json:"assertions"`
	Scripts           []Script    `json:"scripts"`
	Variables         []Variable  `json:"variables"`
}

func (client *Client) ListResults(bucketKey string, testID string) (*[]Result, error) {
	var results = []Result{}

	path := fmt.Sprintf("buckets/%s/tests/%s/results", bucketKey, testID)
	content, err := client.Get(path)
	if err != nil {
		return &results, err
	}

	err = unmarshal(content, &results)
	return &results, err
}

func (client *Client) GetResult(bucketKey string, testID string, testRunID string) (*Result, error) {
	var result = Result{}

	path := fmt.Sprintf("buckets/%s/tests/%s/results/%s", bucketKey, testID, testRunID)
	content, err := client.Get(path)
	if err != nil {
		return &result, err
	}

	err = unmarshal(content, &result)
	return &result, err
}

func (client *Client) GetResultLatest(bucketKey string, testID string) (*Result, error) {
	return client.GetResult(bucketKey, testID, "latest")
}
