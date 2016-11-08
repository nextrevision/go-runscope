package runscope

import (
	"net/http"
	"testing"
)

func TestListIntegrations(t *testing.T) {
	setup()
	defer teardown()

	path := "/teams/1/integrations"
	responseCode := http.StatusOK
	responseData := `
{
  "data": [
    {
      "description": "PagerDuty: Production Alerts",
      "type": "pagerduty",
      "uuid": "cf95026e-8951-4ae1-83a7-699243678490"
    }
  ],
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := []Integration{
		Integration{
			UUID:        "cf95026e-8951-4ae1-83a7-699243678490",
			Type:        "pagerduty",
			Description: "PagerDuty: Production Alerts",
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.ListIntegrations("1")
	if err != nil {
		t.Errorf("ListIntegrations returned error: %v", err)
	}
	testResponseData(t, result, want)
}
