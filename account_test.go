package runscope

import (
	"net/http"
	"testing"
)

func TestGetAccount(t *testing.T) {
	setup()
	defer teardown()

	path := "/account"
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "name": "Grace Hopper",
    "email": "grace@example.com",
    "id": "cf95026e-8951-4ae1-83a7-699243678490",
    "uuid": "cf95026e-8951-4ae1-83a7-699243678490",
    "created_at": 1430512683,
    "teams": [
      {
        "id": "cf95026e-8951-4ae1-83a7-699243678490",
        "name": "TeamAwesome",
        "uuid": "cf95026e-8951-4ae1-83a7-699243678490"
      }
    ]
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := Account{
		Name:      "Grace Hopper",
		Email:     "grace@example.com",
		ID:        "cf95026e-8951-4ae1-83a7-699243678490",
		UUID:      "cf95026e-8951-4ae1-83a7-699243678490",
		CreatedAt: 1430512683,
		Teams: []Team{
			Team{
				Name: "TeamAwesome",
				ID:   "cf95026e-8951-4ae1-83a7-699243678490",
				UUID: "cf95026e-8951-4ae1-83a7-699243678490",
			},
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.GetAccount()
	if err != nil {
		t.Errorf("GetAccount returned error: %v", err)
	}
	testResponseData(t, result, want)
}
