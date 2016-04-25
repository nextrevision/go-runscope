package runscope

import (
	"net/http"
	"testing"
)

func TestListSteps(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/steps"
	responseCode := http.StatusOK
	responseData := `
{
  "data": [
    {
      "id": "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
      "step_type": "request",
      "method": "POST",
      "url": "https://{{base_url}}/example/path",
      "body": "{ \"hello\": \"world\" }",
      "assertions": [
        {
          "source": "response_status",
          "comparison": "equal_number",
          "value": 200
        }
      ],
      "form": {},
      "variables": [
        {
          "name": "source_ip",
          "property": "origin",
          "source": "response_json"
        }
      ],
      "headers": {
        "Content-Type": [
          "application/json"
        ],
        "Accept": [
          "*/*"
        ]
      },
      "scripts": [
        {
          "value": "log(\"This is a sample script\");"
        }
      ],
      "note": "get example data"
    }
  ],
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := &[]Step{
		Step{
			ID:       "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
			StepType: "request",
			Method:   "POST",
			URL:      "https://{{base_url}}/example/path",
			Body:     "{ \"hello\": \"world\" }",
			Form:     make(map[string][]string),
			Assertions: []Assertion{
				Assertion{
					Source:     "response_status",
					Comparison: "equal_number",
					Value:      float64(200),
				},
			},
			Variables: []Variable{
				Variable{
					Name:     "source_ip",
					Property: "origin",
					Source:   "response_json",
				},
			},
			Headers: map[string][]string{
				"Content-Type": []string{
					"application/json",
				},
				"Accept": []string{
					"*/*",
				},
			},
			Scripts: []Script{
				Script{
					Value: "log(\"This is a sample script\");",
				},
			},
			Note: "get example data",
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.ListSteps("1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("ListSteps returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetStep(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/steps/1"
	responseCode := http.StatusOK
	responseData := `
{
	"data": {
    "id": "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
    "step_type": "request",
    "method": "POST",
    "url": "https://{{base_url}}/example/path",
    "body": "{ \"hello\": \"world\" }",
    "assertions": [
      {
        "source": "response_status",
        "comparison": "equal_number",
        "value": 200
      }
    ],
    "form": {},
    "variables": [
      {
        "name": "source_ip",
        "property": "origin",
        "source": "response_json"
      }
    ],
    "headers": {
      "Content-Type": [
        "application/json"
      ],
      "Accept": [
        "*/*"
      ]
    },
    "scripts": [
      {
        "value": "log(\"This is a sample script\");"
      }
    ],
    "note": "get example data"
	},
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := &Step{
		ID:       "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
		StepType: "request",
		Method:   "POST",
		URL:      "https://{{base_url}}/example/path",
		Body:     "{ \"hello\": \"world\" }",
		Form:     make(map[string][]string),
		Assertions: []Assertion{
			Assertion{
				Source:     "response_status",
				Comparison: "equal_number",
				Value:      float64(200),
			},
		},
		Variables: []Variable{
			Variable{
				Name:     "source_ip",
				Property: "origin",
				Source:   "response_json",
			},
		},
		Headers: map[string][]string{
			"Content-Type": []string{
				"application/json",
			},
			"Accept": []string{
				"*/*",
			},
		},
		Scripts: []Script{
			Script{
				Value: "log(\"This is a sample script\");",
			},
		},
		Note: "get example data",
	}

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.GetStep("1", "1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("GetStep returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestNewStep(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/steps"
	request := &Step{
		StepType: "request",
		Method:   "POST",
		URL:      "https://{{base_url}}/example/path",
		Body:     "{ \"hello\": \"world\" }",
		Assertions: []Assertion{
			Assertion{
				Source:     "response_status",
				Comparison: "equal_number",
				Value:      float64(200),
			},
		},
		Variables: []Variable{
			Variable{
				Name:     "source_ip",
				Property: "origin",
				Source:   "response_json",
			},
		},
		Headers: map[string][]string{
			"Content-Type": []string{
				"application/json",
			},
			"Accept": []string{
				"*/*",
			},
		},
		Scripts: []Script{
			Script{
				Value: "log(\"This is a sample script\");",
			},
		},
		Note: "get example data",
	}
	responseCode := http.StatusCreated
	responseData := `
{
  "data": {
    "id": "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
    "step_type": "request",
    "method": "POST",
    "url": "https://{{base_url}}/example/path",
    "body": "{ \"hello\": \"world\" }",
    "assertions": [
      {
        "source": "response_status",
        "comparison": "equal_number",
        "value": 200
      }
    ],
    "form": {},
    "variables": [
      {
        "name": "source_ip",
        "property": "origin",
        "source": "response_json"
      }
    ],
    "headers": {
      "Content-Type": [
        "application/json"
      ],
      "Accept": [
        "*/*"
      ]
    },
    "scripts": [
      {
        "value": "log(\"This is a sample script\");"
      }
    ],
    "note": "get example data"
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := &Step{
		ID:       "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
		StepType: "request",
		Method:   "POST",
		URL:      "https://{{base_url}}/example/path",
		Body:     "{ \"hello\": \"world\" }",
		Form:     make(map[string][]string),
		Assertions: []Assertion{
			Assertion{
				Source:     "response_status",
				Comparison: "equal_number",
				Value:      float64(200),
			},
		},
		Variables: []Variable{
			Variable{
				Name:     "source_ip",
				Property: "origin",
				Source:   "response_json",
			},
		},
		Headers: map[string][]string{
			"Content-Type": []string{
				"application/json",
			},
			"Accept": []string{
				"*/*",
			},
		},
		Scripts: []Script{
			Script{
				Value: "log(\"This is a sample script\");",
			},
		},
		Note: "get example data",
	}

	handlePost(t, path, responseCode, responseData, new(Step), request)

	result, resp, err := client.NewStep("1", "1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("NewStep returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestUpdateStep(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/steps/1"
	request := &Step{
		StepType: "request",
		Method:   "POST",
		URL:      "https://{{base_url}}/example/path",
		Body:     "{ \"hello\": \"world\" }",
		Assertions: []Assertion{
			Assertion{
				Source:     "response_status",
				Comparison: "equal_number",
				Value:      float64(200),
			},
		},
		Variables: []Variable{
			Variable{
				Name:     "source_ip",
				Property: "origin",
				Source:   "response_json",
			},
		},
		Headers: map[string][]string{
			"Content-Type": []string{
				"application/json",
			},
			"Accept": []string{
				"*/*",
			},
		},
		Scripts: []Script{
			Script{
				Value: "log(\"This is a sample script\");",
			},
		},
		Note: "get example data",
	}
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "id": "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
    "step_type": "request",
    "method": "POST",
    "url": "https://{{base_url}}/example/path",
    "body": "{ \"hello\": \"world\" }",
    "assertions": [
      {
        "source": "response_status",
        "comparison": "equal_number",
        "value": 200
      }
    ],
    "form": {},
    "variables": [
      {
        "name": "source_ip",
        "property": "origin",
        "source": "response_json"
      }
    ],
    "headers": {
      "Content-Type": [
        "application/json"
      ],
      "Accept": [
        "*/*"
      ]
    },
    "scripts": [
      {
        "value": "log(\"This is a sample script\");"
      }
    ],
    "note": "get example data"
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := &Step{
		ID:       "b0ecd629-2b92-4kee-9f4e-877c160md9eb",
		StepType: "request",
		Method:   "POST",
		URL:      "https://{{base_url}}/example/path",
		Body:     "{ \"hello\": \"world\" }",
		Form:     make(map[string][]string),
		Assertions: []Assertion{
			Assertion{
				Source:     "response_status",
				Comparison: "equal_number",
				Value:      float64(200),
			},
		},
		Variables: []Variable{
			Variable{
				Name:     "source_ip",
				Property: "origin",
				Source:   "response_json",
			},
		},
		Headers: map[string][]string{
			"Content-Type": []string{
				"application/json",
			},
			"Accept": []string{
				"*/*",
			},
		},
		Scripts: []Script{
			Script{
				Value: "log(\"This is a sample script\");",
			},
		},
		Note: "get example data",
	}

	handlePut(t, path, responseCode, responseData, new(Step), request)

	result, resp, err := client.UpdateStep("1", "1", "1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("UpdateStep returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestDeleteStep(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/steps/1"
	responseCode := http.StatusNoContent

	handleDelete(t, path, responseCode)

	resp, err := client.DeleteStep("1", "1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("DeleteStep returned error: %v", err)
	}
}
