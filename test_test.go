package runscope

import (
	"net/http"
	"testing"
)

func TestListTests(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests"
	responseCode := http.StatusOK
	responseData := `
{
  "data": [
    {
      "created_at": 1438828991,
      "created_by": {
        "email": "grace@example.com",
        "name": "Grace Hopper",
        "id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
      },
      "default_environment_id": "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
      "description": "An internal API!",
      "name": "My Service",
      "id": "9b47981a-98fd-4dac-8f32-c05aa60b8caf"
    }
  ],
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := []Test{
		Test{
			Name:                 "My Service",
			ID:                   "9b47981a-98fd-4dac-8f32-c05aa60b8caf",
			Description:          "An internal API!",
			DefaultEnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
			CreatedAt:            1438828991,
			CreatedBy: Person{
				Name:  "Grace Hopper",
				Email: "grace@example.com",
				ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
			},
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.ListTests("1", ListTestOptions{})
	if err != nil {
		t.Errorf("ListTests returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestListAllTests(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests"
	responseCode := http.StatusOK
	responseData := `
{
  "data": [
    {
      "created_at": 1438828991,
      "created_by": {
        "email": "grace@example.com",
        "name": "Grace Hopper",
        "id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
      },
      "default_environment_id": "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
      "description": "An internal API!",
      "name": "My Service",
      "id": "9b47981a-98fd-4dac-8f32-c05aa60b8caf"
    }
  ],
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := []Test{
		Test{
			Name:                 "My Service",
			ID:                   "9b47981a-98fd-4dac-8f32-c05aa60b8caf",
			Description:          "An internal API!",
			DefaultEnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
			CreatedAt:            1438828991,
			CreatedBy: Person{
				Name:  "Grace Hopper",
				Email: "grace@example.com",
				ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
			},
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.ListAllTests("1")
	if err != nil {
		t.Errorf("ListTests returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetTest(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1"
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "created_at": 1438832081,
    "created_by": {
      "email": "grace@example.com",
      "name": "Grace Hopper",
      "id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
    },
    "default_environment_id": "a50b63cc-c377-4823-9a95-8b91f12326f2",
    "description": null,
    "environments": [
      {
        "emails": {
          "notify_all": false,
          "notify_on": "all",
          "notify_threshold": 1,
          "recipients": []
        },
        "initial_variables": {
          "base_url": "https://api.example.com"
        },
        "integrations": [
          {
            "description": "Pagerduty Account",
            "integration_type": "pagerduty",
            "id": "53776d9a-4f34-4f1f-9gff-c155dfb6692e"
          }
        ],
        "name": "Test Settings",
        "parent_environment_id": null,
        "preserve_cookies": false,
        "regions": [
          "us1"
        ],
        "remote_agents": [],
        "script": "",
        "test_id": "626a024c-f75e-4f57-82d4-104fe443c0f3",
        "id": "a50b63cc-c377-4823-9a95-8b91f12326f2",
        "verify_ssl": true,
        "webhooks": null
      }
    ],
    "last_run": null,
    "name": "Sample Name",
    "schedules": [],
    "steps": [
      {
        "assertions": [
          {
            "comparison": "is_equal",
            "source": "response_status",
            "value": 200
          }
        ],
        "auth": {},
        "body": "",
        "form": {},
        "headers": {},
        "method": "GET",
        "note": "",
        "step_type": "request",
        "url": "https://yourapihere.com/",
        "id": "53f8e1fd-0989-491a-9f15-cc055f27d097",
        "variables": []
      }
    ],
    "trigger_url": "http://api.runscope.com/radar/b96ecee2-cce6-4d80-8f07-33ac22a22ebd/trigger",
    "id": "626a024c-f75e-4f57-82d4-104fe443c0f3"
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := Test{
		Name: "Sample Name",
		ID:   "626a024c-f75e-4f57-82d4-104fe443c0f3",
		CreatedBy: Person{
			Name:  "Grace Hopper",
			Email: "grace@example.com",
			ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
		},
		CreatedAt:            1438832081,
		DefaultEnvironmentID: "a50b63cc-c377-4823-9a95-8b91f12326f2",
		TriggerURL:           "http://api.runscope.com/radar/b96ecee2-cce6-4d80-8f07-33ac22a22ebd/trigger",
		Steps: []Step{
			Step{
				StepType: "request",
				ID:       "53f8e1fd-0989-491a-9f15-cc055f27d097",
				Method:   "GET",
				URL:      "https://yourapihere.com/",
				Assertions: []Assertion{
					Assertion{
						Comparison: "is_equal",
						Source:     "response_status",
						Value:      float64(200),
					},
				},
				Variables: []Variable{},
				Headers:   make(map[string][]string),
				Form:      make(map[string][]string),
			},
		},
		Schedules: []Schedule{},
		Environments: []Environment{
			Environment{
				Name: "Test Settings",
				ID:   "a50b63cc-c377-4823-9a95-8b91f12326f2",
				Regions: []string{
					"us1",
				},
				RemoteAgents: []RemoteAgent{},
				VerifySSL:    true,
				TestID:       "626a024c-f75e-4f57-82d4-104fe443c0f3",
				Emails: Email{
					NotifyOn:        "all",
					NotifyThreshold: 1,
					Recipients:      []Person{},
				},
				InitialVariables: map[string]string{
					"base_url": "https://api.example.com",
				},
				Integrations: []TeamIntegration{
					TeamIntegration{
						ID:          "53776d9a-4f34-4f1f-9gff-c155dfb6692e",
						Description: "Pagerduty Account",
						Type:        "pagerduty",
					},
				},
			},
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.GetTest("1", "1")
	if err != nil {
		t.Errorf("GetTest returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestNewTest(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests"
	request := NewTestRequest{
		Name:        "Sample Test",
		Description: "A new sample test",
	}
	responseCode := http.StatusCreated
	responseData := `
{
  "data": {
    "name": "Sample Test",
    "Description": "A new sample test"
  },
  "meta": {
    "status": "success"
  }
}`
	want := Test{
		Name:        "Sample Test",
		Description: "A new sample test",
	}

	handlePost(t, path, responseCode, responseData, new(NewTestRequest), request)

	result, err := client.NewTest("1", request)
	if err != nil {
		t.Errorf("NewTest returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestUpdateTest(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1"
	request := UpdateTestRequest{
		Name:        "Sample Test",
		Description: "A new sample test",
	}
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "name": "Sample Test",
    "Description": "A new sample test"
  },
  "meta": {
    "status": "success"
  }
}`
	want := Test{
		Name:        "Sample Test",
		Description: "A new sample test",
	}

	handlePut(t, path, responseCode, responseData, new(UpdateTestRequest), request)

	result, err := client.UpdateTest("1", "1", request)
	if err != nil {
		t.Errorf("UpdateTest returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestDeleteTest(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1"
	responseCode := http.StatusNoContent

	handleDelete(t, path, responseCode)

	err := client.DeleteTest("1", "1")
	if err != nil {
		t.Errorf("DeleteTest returned error: %v", err)
	}
}
