package runscope

import (
	"net/http"
	"testing"
)

func TestListResults(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/results"
	responseCode := http.StatusOK
	responseData := `
{
  "data": [
    {
      "agent": null,
      "assertions_defined": 2,
      "assertions_failed": 0,
      "assertions_passed": 2,
      "bucket_key": "6knqwmwvqpzr",
      "finished_at": 1406061608.506811,
      "region": "us1",
      "requests_executed": 1,
      "result": "pass",
      "scripts_defined": 2,
      "scripts_failed": 0,
      "scripts_passed": 2,
      "started_at": 1406036406.68105,
      "test_run_id": "0aa48464-f89e-4596-8d60-79bc678d313f",
      "test_run_url": "https://api.runscope.com/buckets/6knqwmwvqpzr/tests/db4cc896-2804-4520-ad06-0caf3bf216a8/results/0aa48464-f89e-4596-8d60-79bc678d313f",
      "test_id": "db4cc896-2804-4520-ad06-0caf3bf216a8",
      "variables_defined": 2,
      "variables_failed": 0,
      "variables_passed": 2,
      "environment_id": "abcdc896-2804-4520-ad06-0caf3bf216a8",
      "environment_name": "My Test Environment"
    }
  ],
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := []Result{
		Result{
			AssertionsDefined: 2,
			AssertionsFailed:  0,
			AssertionsPassed:  2,
			BucketKey:         "6knqwmwvqpzr",
			FinishedAt:        1406061608.506811,
			Region:            "us1",
			RequestsExecuted:  1,
			Result:            "pass",
			ScriptsDefined:    2,
			ScriptsFailed:     0,
			ScriptsPassed:     2,
			StartedAt:         1406036406.68105,
			TestRunID:         "0aa48464-f89e-4596-8d60-79bc678d313f",
			TestRunURL:        "https://api.runscope.com/buckets/6knqwmwvqpzr/tests/db4cc896-2804-4520-ad06-0caf3bf216a8/results/0aa48464-f89e-4596-8d60-79bc678d313f",
			TestID:            "db4cc896-2804-4520-ad06-0caf3bf216a8",
			VariablesDefined:  2,
			VariablesFailed:   0,
			VariablesPassed:   2,
			EnvironmentID:     "abcdc896-2804-4520-ad06-0caf3bf216a8",
			EnvironmentName:   "My Test Environment",
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.ListResults("1", "1")
	if err != nil {
		t.Errorf("ListResults returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetResult(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/results/1"
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "agent": null,
    "assertions_defined": 2,
    "assertions_failed": 0,
    "assertions_passed": 2,
    "bucket_key": "6knqwmwvqpzr",
    "finished_at": 1406061608.506811,
    "region": "us1",
    "requests": [
      {
        "result": "pass",
        "url": "https://yourapihere.com/",
        "method": "GET",
        "assertions_defined": 1,
        "assertions_failed": 0,
        "assertions_passed": 1,
        "scripts_defined": 1,
        "scripts_failed": 0,
        "scripts_passed": 1,
        "variables_defined": 1,
        "variables_failed": 0,
        "variables_passed": 1,
        "assertions": [
          {
            "result": "pass",
            "source": "json",
            "property": "id",
            "comparison": "equals_number",
            "target_value": "123",
            "actual_value": "123",
            "error": null
          }
        ],
        "scripts": [
          {
            "result": "pass",
            "output": "script output",
            "error": null
          }
        ],
        "variables": [
          {
            "result": "pass",
            "source": "json",
            "property": "id",
            "name": "customer_id",
            "value": 123,
            "error": null
          }
        ]
      }
    ],
    "requests_executed": 1,
    "result": "pass",
    "scripts_defined": 2,
    "scripts_failed": 0,
    "scripts_passed": 2,
    "started_at": 1406036406.68105,
    "test_run_id": "0aa48464-f89e-4596-8d60-79bc678d313f",
    "test_run_url": "https://api.runscope.com/buckets/6knqwmwvqpzr/tests/db4cc896-2804-4520-ad06-0caf3bf216a8/results/0aa48464-f89e-4596-8d60-79bc678d313f",
    "test_id": "db4cc896-2804-4520-ad06-0caf3bf216a8",
    "variables_defined": 2,
    "variables_failed": 0,
    "variables_passed": 2,
    "environment_id": "abcdc896-2804-4520-ad06-0caf3bf216a8",
    "environment_name": "My Test Environment"
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := Result{
		AssertionsDefined: 2,
		AssertionsFailed:  0,
		AssertionsPassed:  2,
		BucketKey:         "6knqwmwvqpzr",
		FinishedAt:        1406061608.506811,
		Region:            "us1",
		RequestsExecuted:  1,
		Result:            "pass",
		ScriptsDefined:    2,
		ScriptsFailed:     0,
		ScriptsPassed:     2,
		StartedAt:         1406036406.68105,
		TestRunID:         "0aa48464-f89e-4596-8d60-79bc678d313f",
		TestRunURL:        "https://api.runscope.com/buckets/6knqwmwvqpzr/tests/db4cc896-2804-4520-ad06-0caf3bf216a8/results/0aa48464-f89e-4596-8d60-79bc678d313f",
		TestID:            "db4cc896-2804-4520-ad06-0caf3bf216a8",
		VariablesDefined:  2,
		VariablesFailed:   0,
		VariablesPassed:   2,
		EnvironmentID:     "abcdc896-2804-4520-ad06-0caf3bf216a8",
		EnvironmentName:   "My Test Environment",
		Requests: []Request{
			Request{
				Result:            "pass",
				URL:               "https://yourapihere.com/",
				Method:            "GET",
				AssertionsDefined: 1,
				AssertionsFailed:  0,
				AssertionsPassed:  1,
				ScriptsDefined:    1,
				ScriptsFailed:     0,
				ScriptsPassed:     1,
				VariablesDefined:  1,
				VariablesFailed:   0,
				VariablesPassed:   1,
				Assertions: []Assertion{
					Assertion{
						Result:      "pass",
						Source:      "json",
						Property:    "id",
						Comparison:  "equals_number",
						TargetValue: "123",
						ActualValue: "123",
					},
				},
				Scripts: []Script{
					Script{
						Result: "pass",
						Output: "script output",
					},
				},
				Variables: []Variable{
					Variable{
						Result:   "pass",
						Source:   "json",
						Property: "id",
						Name:     "customer_id",
						Value:    float64(123),
					},
				},
			},
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.GetResult("1", "1", "1")
	if err != nil {
		t.Errorf("GetResult returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetResultLatest(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/results/latest"
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "agent": null,
    "assertions_defined": 2,
    "assertions_failed": 0,
    "assertions_passed": 2,
    "bucket_key": "6knqwmwvqpzr",
    "finished_at": 1406061608.506811,
    "region": "us1",
    "requests": [
      {
        "result": "pass",
        "url": "https://yourapihere.com/",
        "method": "GET",
        "assertions_defined": 1,
        "assertions_failed": 0,
        "assertions_passed": 1,
        "scripts_defined": 1,
        "scripts_failed": 0,
        "scripts_passed": 1,
        "variables_defined": 1,
        "variables_failed": 0,
        "variables_passed": 1,
        "assertions": [
          {
            "result": "pass",
            "source": "json",
            "property": "id",
            "comparison": "equals_number",
            "target_value": "123",
            "actual_value": "123",
            "error": null
          }
        ],
        "scripts": [
          {
            "result": "pass",
            "output": "script output",
            "error": null
          }
        ],
        "variables": [
          {
            "result": "pass",
            "source": "json",
            "property": "id",
            "name": "customer_id",
            "value": 123,
            "error": null
          }
        ]
      }
    ],
    "requests_executed": 1,
    "result": "pass",
    "scripts_defined": 2,
    "scripts_failed": 0,
    "scripts_passed": 2,
    "started_at": 1406036406.68105,
    "test_run_id": "0aa48464-f89e-4596-8d60-79bc678d313f",
    "test_run_url": "https://api.runscope.com/buckets/6knqwmwvqpzr/tests/db4cc896-2804-4520-ad06-0caf3bf216a8/results/0aa48464-f89e-4596-8d60-79bc678d313f",
    "test_id": "db4cc896-2804-4520-ad06-0caf3bf216a8",
    "variables_defined": 2,
    "variables_failed": 0,
    "variables_passed": 2,
    "environment_id": "abcdc896-2804-4520-ad06-0caf3bf216a8",
    "environment_name": "My Test Environment"
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := Result{
		AssertionsDefined: 2,
		AssertionsFailed:  0,
		AssertionsPassed:  2,
		BucketKey:         "6knqwmwvqpzr",
		FinishedAt:        1406061608.506811,
		Region:            "us1",
		RequestsExecuted:  1,
		Result:            "pass",
		ScriptsDefined:    2,
		ScriptsFailed:     0,
		ScriptsPassed:     2,
		StartedAt:         1406036406.68105,
		TestRunID:         "0aa48464-f89e-4596-8d60-79bc678d313f",
		TestRunURL:        "https://api.runscope.com/buckets/6knqwmwvqpzr/tests/db4cc896-2804-4520-ad06-0caf3bf216a8/results/0aa48464-f89e-4596-8d60-79bc678d313f",
		TestID:            "db4cc896-2804-4520-ad06-0caf3bf216a8",
		VariablesDefined:  2,
		VariablesFailed:   0,
		VariablesPassed:   2,
		EnvironmentID:     "abcdc896-2804-4520-ad06-0caf3bf216a8",
		EnvironmentName:   "My Test Environment",
		Requests: []Request{
			Request{
				Result:            "pass",
				URL:               "https://yourapihere.com/",
				Method:            "GET",
				AssertionsDefined: 1,
				AssertionsFailed:  0,
				AssertionsPassed:  1,
				ScriptsDefined:    1,
				ScriptsFailed:     0,
				ScriptsPassed:     1,
				VariablesDefined:  1,
				VariablesFailed:   0,
				VariablesPassed:   1,
				Assertions: []Assertion{
					Assertion{
						Result:      "pass",
						Source:      "json",
						Property:    "id",
						Comparison:  "equals_number",
						TargetValue: "123",
						ActualValue: "123",
					},
				},
				Scripts: []Script{
					Script{
						Result: "pass",
						Output: "script output",
					},
				},
				Variables: []Variable{
					Variable{
						Result:   "pass",
						Source:   "json",
						Property: "id",
						Name:     "customer_id",
						Value:    float64(123),
					},
				},
			},
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.GetResultLatest("1", "1")
	if err != nil {
		t.Errorf("GetResultLatest returned error: %v", err)
	}
	testResponseData(t, result, want)
}
