package runscope

import (
	"net/http"
	"testing"
)

func TestListTestEnvironments(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/environments"
	responseCode := http.StatusOK
	responseData := `
{
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
}`
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

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.ListTestEnvironments("1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("ListTestEnvironments returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestListSharedEnvironments(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/environments"
	responseCode := http.StatusOK
	responseData := `
{
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
}`
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

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.ListSharedEnvironments("1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("ListSharedEnvironments returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetTestEnvironment(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/environments/1"
	responseCode := http.StatusOK
	responseData := `
{
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
}`
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

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.GetTestEnvironment("1", "1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("GetTestEnvironment returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetSharedEnvironment(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/environments/1"
	responseCode := http.StatusOK
	responseData := `
{
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
}`
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

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.GetSharedEnvironment("1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("GetSharedEnvironment returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestNewTestEnvironment(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/environments"
	request := &Environment{
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
	responseCode := http.StatusCreated
	responseData := `
{
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
}`
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

	handlePost(t, path, responseCode, responseData, new(Environment), request)

	result, resp, err := client.NewTestEnvironment("1", "1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("NewTestEnvironment returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestNewSharedEnvironment(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/environments"
	request := &Environment{
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
	responseCode := http.StatusCreated
	responseData := `
{
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
}`
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

	handlePost(t, path, responseCode, responseData, new(Environment), request)

	result, resp, err := client.NewSharedEnvironment("1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("NewTestEnvironment returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestUpdateTestEnvironment(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/environments/1"
	request := &Environment{
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
	responseCode := http.StatusOK
	responseData := `
{
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
}`
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

	handlePut(t, path, responseCode, responseData, new(Environment), request)

	result, resp, err := client.UpdateTestEnvironment("1", "1", "1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("UpdateTestEnvironment returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestUpdateSharedEnvironment(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/environments/1"
	request := &Environment{
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
	responseCode := http.StatusOK
	responseData := `
{
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
}`
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

	handlePut(t, path, responseCode, responseData, new(Environment), request)

	result, resp, err := client.UpdateSharedEnvironment("1", "1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("UpdateSharedEnvironment returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestDeleteEnvironment(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/environments/1"
	responseCode := http.StatusNoContent

	handleDelete(t, path, responseCode)

	resp, err := client.DeleteEnvironment("1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("DeleteEnvironment returned error: %v", err)
	}
}
