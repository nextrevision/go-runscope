package runscope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListSteps(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/tests/1/steps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
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
			}`)
	})

	steps, resp, err := client.ListSteps("1", "1")
	if err != nil {
		t.Errorf("ListSteps returned error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("ListSteps did not return 200: %v", resp)
	}

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
	if !reflect.DeepEqual(steps, want) {
		t.Errorf("ListSteps returned %+v, want %+v", steps, want)
	}
}

func TestGetStep(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/tests/1/steps/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
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
			}`)
	})

	step, resp, err := client.GetStep("1", "1", "1")
	if err != nil {
		t.Errorf("GetStep returned error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("GetStep did not return 200: %v", resp)
	}

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
	if !reflect.DeepEqual(step, want) {
		t.Errorf("GetStep returned %+v, want %+v", step, want)
	}
}

func TestNewStep(t *testing.T) {
	setup()
	defer teardown()

	req := &Step{
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

	mux.HandleFunc("/buckets/1/tests/1/steps", func(w http.ResponseWriter, r *http.Request) {
		v := new(Step)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,
			`{
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
			}`)
	})

	step, resp, err := client.NewStep("1", "1", req)
	if err != nil {
		t.Errorf("NewStep returned error: %v", err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("NewStep did not return 201: %v", resp)
	}

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
	if !reflect.DeepEqual(step, want) {
		t.Errorf("NewStep returned %+v, want %+v", step, want)
	}
}

func TestUpdateStep(t *testing.T) {
	setup()
	defer teardown()

	req := &Step{
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

	mux.HandleFunc("/buckets/1/tests/1/steps/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(Step)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "PUT")
		fmt.Fprint(w,
			`{
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
			}`)
	})

	step, resp, err := client.UpdateStep("1", "1", "1", req)
	if err != nil {
		t.Errorf("UpdateStep returned error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("UpdateStep did not return 200: %v", resp)
	}

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
	if !reflect.DeepEqual(step, want) {
		t.Errorf("UpdateStep returned %+v, want %+v", step, want)
	}
}

func TestDeleteStep(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/tests/1/steps/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.DeleteStep("1", "1", "1")
	if err != nil {
		t.Errorf("DeleteStep returned error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("DeleteStep did not return 204: %v", resp)
	}
}
