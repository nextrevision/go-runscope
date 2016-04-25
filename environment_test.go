package runscope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListTestEnvironments(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/tests/1/environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"data": [
					{
						"emails": {
							"notify_all": false,
							"notify_on": "all",
							"notify_threshold": null,
							"recipients": [
								{
									"email": "grace@example.com",
									"name": "Grace Hopper",
									"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
								}
							]
						},
						"initial_variables": {
							"my_variable": "some value",
							"one more": "values"
						},
						"integrations": [],
						"name": "Remote Settings",
						"parent_environment_id": null,
						"preserve_cookies": false,
						"regions": [
							"us1",
							"jp1"
						],
						"remote_agents": [
							{
								"name": "my-local-machine.runscope.com",
								"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
							}
						],
						"script": "var a = \"asdf\";\nlog(\"OK\");",
						"test_id": null,
						"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
						"verify_ssl": true,
						"webhooks": [
							"http://api.example.com/webhook_reciever",
							"https://yourapihere.com/post"
						]
					}
				],
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environments, _, err := client.ListTestEnvironments("1", "1")
	if err != nil {
		t.Errorf("ListTestEnvironments returned error: %v", err)
	}

	want := &[]Environment{
		Environment{
			Name:            "Remote Settings",
			ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
			PreserveCookies: false,
			Regions:         []string{"us1", "jp1"},
			RemoteAgents: []RemoteAgent{
				RemoteAgent{
					Name: "my-local-machine.runscope.com",
					UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
				},
			},
			Script:    "var a = \"asdf\";\nlog(\"OK\");",
			VerifySSL: true,
			Webhooks: []string{
				"http://api.example.com/webhook_reciever",
				"https://yourapihere.com/post",
			},
			Emails: Email{
				NotifyAll: false,
				NotifyOn:  "all",
				Recipients: []Person{
					Person{
						Name:  "Grace Hopper",
						Email: "grace@example.com",
						ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
					},
				},
			},
			InitialVariables: map[string]string{
				"my_variable": "some value",
				"one more":    "values",
			},
			Integrations: make([]Integration, 0),
		},
	}
	if !reflect.DeepEqual(environments, want) {
		t.Errorf("ListTestEnvironments returned %+v, want %+v", environments, want)
	}
}

func TestListSharedEnvironments(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"data": [
					{
						"emails": {
							"notify_all": false,
							"notify_on": "all",
							"notify_threshold": null,
							"recipients": [
								{
									"email": "grace@example.com",
									"name": "Grace Hopper",
									"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
								}
							]
						},
						"initial_variables": {
							"my_variable": "some value",
							"one more": "values"
						},
						"integrations": [],
						"name": "Remote Settings",
						"parent_environment_id": null,
						"preserve_cookies": false,
						"regions": [
							"us1",
							"jp1"
						],
						"remote_agents": [
							{
								"name": "my-local-machine.runscope.com",
								"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
							}
						],
						"script": "var a = \"asdf\";\nlog(\"OK\");",
						"test_id": null,
						"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
						"verify_ssl": true,
						"webhooks": [
							"http://api.example.com/webhook_reciever",
							"https://yourapihere.com/post"
						]
					}
				],
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environments, _, err := client.ListSharedEnvironments("1")
	if err != nil {
		t.Errorf("ListSharedEnvironments returned error: %v", err)
	}

	want := &[]Environment{
		Environment{
			Name:            "Remote Settings",
			ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
			PreserveCookies: false,
			Regions:         []string{"us1", "jp1"},
			RemoteAgents: []RemoteAgent{
				RemoteAgent{
					Name: "my-local-machine.runscope.com",
					UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
				},
			},
			Script:    "var a = \"asdf\";\nlog(\"OK\");",
			VerifySSL: true,
			Webhooks: []string{
				"http://api.example.com/webhook_reciever",
				"https://yourapihere.com/post",
			},
			Emails: Email{
				NotifyAll: false,
				NotifyOn:  "all",
				Recipients: []Person{
					Person{
						Name:  "Grace Hopper",
						Email: "grace@example.com",
						ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
					},
				},
			},
			InitialVariables: map[string]string{
				"my_variable": "some value",
				"one more":    "values",
			},
			Integrations: make([]Integration, 0),
		},
	}
	if !reflect.DeepEqual(environments, want) {
		t.Errorf("ListSharedEnvironments returned %+v, want %+v", environments, want)
	}
}

func TestGetTestEnvironment(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/tests/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"data": {
					"emails": {
						"notify_all": false,
						"notify_on": "all",
						"notify_threshold": null,
						"recipients": [
							{
								"email": "grace@example.com",
								"name": "Grace Hopper",
								"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
							}
						]
					},
					"initial_variables": {
						"my_variable": "some value",
						"one more": "values"
					},
					"integrations": [],
					"name": "Remote Settings",
					"parent_environment_id": null,
					"preserve_cookies": false,
					"regions": [
						"us1",
						"jp1"
					],
					"remote_agents": [
						{
							"name": "my-local-machine.runscope.com",
							"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
						}
					],
					"script": "var a = \"asdf\";\nlog(\"OK\");",
					"test_id": null,
					"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
					"verify_ssl": true,
					"webhooks": [
						"http://api.example.com/webhook_reciever",
						"https://yourapihere.com/post"
					]
				},
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environment, _, err := client.GetTestEnvironment("1", "1", "1")
	if err != nil {
		t.Errorf("GetTestEnvironment returned error: %v", err)
	}

	want := &Environment{
		Name:            "Remote Settings",
		ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
		PreserveCookies: false,
		Regions:         []string{"us1", "jp1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		Script:    "var a = \"asdf\";\nlog(\"OK\");",
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
			"one more":    "values",
		},
		Integrations: make([]Integration, 0),
	}
	if !reflect.DeepEqual(environment, want) {
		t.Errorf("GetTestEnvironment returned %+v, want %+v", environment, want)
	}
}

func TestGetSharedEnvironment(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"data": {
					"emails": {
						"notify_all": false,
						"notify_on": "all",
						"notify_threshold": null,
						"recipients": [
							{
								"email": "grace@example.com",
								"name": "Grace Hopper",
								"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
							}
						]
					},
					"initial_variables": {
						"my_variable": "some value",
						"one more": "values"
					},
					"integrations": [],
					"name": "Remote Settings",
					"parent_environment_id": null,
					"preserve_cookies": false,
					"regions": [
						"us1",
						"jp1"
					],
					"remote_agents": [
						{
							"name": "my-local-machine.runscope.com",
							"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
						}
					],
					"script": "var a = \"asdf\";\nlog(\"OK\");",
					"test_id": null,
					"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
					"verify_ssl": true,
					"webhooks": [
						"http://api.example.com/webhook_reciever",
						"https://yourapihere.com/post"
					]
				},
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environment, _, err := client.GetSharedEnvironment("1", "1")
	if err != nil {
		t.Errorf("GetSharedEnvironment returned error: %v", err)
	}

	want := &Environment{
		Name:            "Remote Settings",
		ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
		PreserveCookies: false,
		Regions:         []string{"us1", "jp1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		Script:    "var a = \"asdf\";\nlog(\"OK\");",
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
			"one more":    "values",
		},
		Integrations: make([]Integration, 0),
	}
	if !reflect.DeepEqual(environment, want) {
		t.Errorf("GetSharedEnvironment returned %+v, want %+v", environment, want)
	}
}

func TestNewTestEnvironment(t *testing.T) {
	setup()
	defer teardown()

	req := &Environment{
		Name:    "Remote Settings",
		Regions: []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
	}

	mux.HandleFunc("/buckets/1/tests/1/environments", func(w http.ResponseWriter, r *http.Request) {
		v := new(Environment)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,
			`{
				"data": {
					"emails": {
						"notify_all": false,
						"notify_on": "all",
						"notify_threshold": null,
						"recipients": [
							{
								"email": "grace@example.com",
								"name": "Grace Hopper",
								"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
							}
						]
					},
					"initial_variables": {
						"my_variable": "some value"
					},
					"integrations": [],
					"name": "Remote Settings",
					"parent_environment_id": null,
					"preserve_cookies": false,
					"regions": [
						"us1"
					],
					"remote_agents": [
						{
							"name": "my-local-machine.runscope.com",
							"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
						}
					],
					"script": null,
					"test_id": null,
					"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
					"verify_ssl": true,
					"webhooks": [
						"http://api.example.com/webhook_reciever",
						"https://yourapihere.com/post"
					]
				},
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environment, resp, err := client.NewTestEnvironment("1", "1", req)
	if err != nil {
		t.Errorf("NewTestEnvironment returned error: %v", err)
	}

	want := &Environment{
		Name:            "Remote Settings",
		ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
		PreserveCookies: false,
		Regions:         []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
		Integrations: make([]Integration, 0),
	}
	if resp.StatusCode != 201 {
		t.Errorf("NewTestEnvironment did not return 201: %v", resp)
	}
	if !reflect.DeepEqual(environment, want) {
		t.Errorf("NewTestEnvironment returned %+v, want %+v", environment, want)
	}
}

func TestNewSharedEnvironment(t *testing.T) {
	setup()
	defer teardown()

	req := &Environment{
		Name:    "Remote Settings",
		Regions: []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
	}

	mux.HandleFunc("/buckets/1/environments", func(w http.ResponseWriter, r *http.Request) {
		v := new(Environment)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,
			`{
				"data": {
					"emails": {
						"notify_all": false,
						"notify_on": "all",
						"notify_threshold": null,
						"recipients": [
							{
								"email": "grace@example.com",
								"name": "Grace Hopper",
								"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
							}
						]
					},
					"initial_variables": {
						"my_variable": "some value"
					},
					"integrations": [],
					"name": "Remote Settings",
					"parent_environment_id": null,
					"preserve_cookies": false,
					"regions": [
						"us1"
					],
					"remote_agents": [
						{
							"name": "my-local-machine.runscope.com",
							"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
						}
					],
					"script": null,
					"test_id": null,
					"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
					"verify_ssl": true,
					"webhooks": [
						"http://api.example.com/webhook_reciever",
						"https://yourapihere.com/post"
					]
				},
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environment, resp, err := client.NewSharedEnvironment("1", req)
	if err != nil {
		t.Errorf("NewSharedEnvironment returned error: %v", err)
	}

	want := &Environment{
		Name:            "Remote Settings",
		ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
		PreserveCookies: false,
		Regions:         []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
		Integrations: make([]Integration, 0),
	}
	if resp.StatusCode != 201 {
		t.Errorf("NewSharedEnvironment did not return 201: %v", resp)
	}
	if !reflect.DeepEqual(environment, want) {
		t.Errorf("NewSharedEnvironment returned %+v, want %+v", environment, want)
	}
}

func TestUpdateTestEnvironment(t *testing.T) {
	setup()
	defer teardown()

	req := &Environment{
		Name:    "Remote Settings",
		ID:      "f8007150-0052-482c-9d52-c3ea4042e0f5",
		Regions: []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
	}

	mux.HandleFunc("/buckets/1/tests/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(Environment)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
				"data": {
					"emails": {
						"notify_all": false,
						"notify_on": "all",
						"notify_threshold": null,
						"recipients": [
							{
								"email": "grace@example.com",
								"name": "Grace Hopper",
								"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
							}
						]
					},
					"initial_variables": {
						"my_variable": "some value"
					},
					"integrations": [],
					"name": "Remote Settings",
					"parent_environment_id": null,
					"preserve_cookies": false,
					"regions": [
						"us1"
					],
					"remote_agents": [
						{
							"name": "my-local-machine.runscope.com",
							"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
						}
					],
					"script": null,
					"test_id": null,
					"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
					"verify_ssl": true,
					"webhooks": [
						"http://api.example.com/webhook_reciever",
						"https://yourapihere.com/post"
					]
				},
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environment, _, err := client.UpdateTestEnvironment("1", "1", "1", req)
	if err != nil {
		t.Errorf("UpdateTestEnvironment returned error: %v", err)
	}

	want := &Environment{
		Name:            "Remote Settings",
		ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
		PreserveCookies: false,
		Regions:         []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
		Integrations: make([]Integration, 0),
	}
	if !reflect.DeepEqual(environment, want) {
		t.Errorf("UpdateTestEnvironment returned %+v, want %+v", environment, want)
	}
}

func TestUpdateSharedEnvironment(t *testing.T) {
	setup()
	defer teardown()

	req := &Environment{
		Name:    "Remote Settings",
		ID:      "f8007150-0052-482c-9d52-c3ea4042e0f5",
		Regions: []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
	}

	mux.HandleFunc("/buckets/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(Environment)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
				"data": {
					"emails": {
						"notify_all": false,
						"notify_on": "all",
						"notify_threshold": null,
						"recipients": [
							{
								"email": "grace@example.com",
								"name": "Grace Hopper",
								"id": "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9"
							}
						]
					},
					"initial_variables": {
						"my_variable": "some value"
					},
					"integrations": [],
					"name": "Remote Settings",
					"parent_environment_id": null,
					"preserve_cookies": false,
					"regions": [
						"us1"
					],
					"remote_agents": [
						{
							"name": "my-local-machine.runscope.com",
							"uuid": "141d4dbc-1e41-401e-8067-6df18501e9ed"
						}
					],
					"script": null,
					"test_id": null,
					"id": "f8007150-0052-482c-9d52-c3ea4042e0f5",
					"verify_ssl": true,
					"webhooks": [
						"http://api.example.com/webhook_reciever",
						"https://yourapihere.com/post"
					]
				},
				"error": null,
				"meta": {
					"status": "success"
				}
			}`)
	})

	environment, _, err := client.UpdateSharedEnvironment("1", "1", req)
	if err != nil {
		t.Errorf("UpdateSharedEnvironment returned error: %v", err)
	}

	want := &Environment{
		Name:            "Remote Settings",
		ID:              "f8007150-0052-482c-9d52-c3ea4042e0f5",
		PreserveCookies: false,
		Regions:         []string{"us1"},
		RemoteAgents: []RemoteAgent{
			RemoteAgent{
				Name: "my-local-machine.runscope.com",
				UUID: "141d4dbc-1e41-401e-8067-6df18501e9ed",
			},
		},
		VerifySSL: true,
		Webhooks: []string{
			"http://api.example.com/webhook_reciever",
			"https://yourapihere.com/post",
		},
		Emails: Email{
			NotifyAll: false,
			NotifyOn:  "all",
			Recipients: []Person{
				Person{
					Name:  "Grace Hopper",
					Email: "grace@example.com",
					ID:    "4ee15ecc-7fe1-43cb-aa12-ef50420f2cf9",
				},
			},
		},
		InitialVariables: map[string]string{
			"my_variable": "some value",
		},
		Integrations: make([]Integration, 0),
	}
	if !reflect.DeepEqual(environment, want) {
		t.Errorf("UpdateSharedEnvironment returned %+v, want %+v", environment, want)
	}
}

func TestDeleteEnvironment(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.DeleteEnvironment("1", "1")
	if resp.StatusCode != 204 {
		t.Errorf("DeleteEnvironment did not return 204: %v", resp)
	}
	if err != nil {
		t.Errorf("DeleteEnvironment returned error: %v", err)
	}
}
