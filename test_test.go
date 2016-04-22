package runscope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestListTests(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345/tests",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w,
				`{
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
				}`)
		},
	)

	tests, _, err := client.ListTests("abcde12345")
	if err != nil {
		t.Errorf("ListTests returned error: %v", err)
	}

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
	if !reflect.DeepEqual(tests, want) {
		t.Errorf("ListTests returned %+v, want %+v",
			tests, want)
	}
}

// TODO: return proper http response code
func TestGetTest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345/tests/12345",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w,
				`{
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
        }`)
		},
	)

	test, _, err := client.GetTest("abcde12345", "12345")
	if err != nil {
		t.Errorf("GetTest returned error: %v", err)
	}

	want := &Test{
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
				Integrations: []Integration{
					Integration{
						ID:          "53776d9a-4f34-4f1f-9gff-c155dfb6692e",
						Description: "Pagerduty Account",
						Type:        "pagerduty",
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(test, want) {
		fmt.Printf("%# v", pretty.Formatter(test))
		fmt.Printf("%# v", pretty.Formatter(want))
		t.Errorf("GetTest returned %+v, want %+v",
			test, want)
	}
}

func TestNewTest(t *testing.T) {
	setup()
	defer teardown()

	input := &Test{
		Name:        "Sample Test",
		Description: "A new sample test",
	}

	mux.HandleFunc("/buckets/abcde12345/tests",
		func(w http.ResponseWriter, r *http.Request) {
			v := new(Test)
			json.NewDecoder(r.Body).Decode(v)

			if !reflect.DeepEqual(v, input) {
				t.Errorf("Request body = %+v, want %+v", v, input)
			}
			testMethod(t, r, "POST")
			fmt.Fprint(w,
				`{
           "data": {
						  "name": "Sample Test",
							"Description": "A new sample test"
          },
          "meta": {
            "status": "success"
          }
        }`)
		},
	)

	test, _, err := client.NewTest("abcde12345", input)
	if err != nil {
		t.Errorf("NewTest returned error: %v", err)
	}

	want := &Test{
		Name:        "Sample Test",
		Description: "A new sample test",
	}
	if !reflect.DeepEqual(test, want) {
		t.Errorf("NewTest returned %+v, want %+v",
			test, want)
	}
}

func TestDeleteTest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345/tests/12345",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			w.WriteHeader(http.StatusNoContent)
		},
	)

	resp, err := client.DeleteTest("abcde12345", "12345")
	if resp.StatusCode != 204 {
		t.Errorf("DeleteTest did not return 204: %v", resp)
	}
	if err != nil {
		t.Errorf("DeleteTest returned error: %v", err)
	}
}
