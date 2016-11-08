package runscope

import (
	"net/http"
	"testing"
)

func TestListRegions(t *testing.T) {
	setup()
	defer teardown()

	path := "/regions"
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "regions": [
      {
        "hostname": "us1.runscope.net",
        "region_code": "us1",
        "location": "US Virginia (None)",
        "service_provider": "Amazon Web Services"
      }
    ]
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := Regions{
		Regions: []Region{
			Region{
				Hostname:        "us1.runscope.net",
				RegionCode:      "us1",
				Location:        "US Virginia (None)",
				ServiceProvider: "Amazon Web Services",
			},
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.ListRegions()
	if err != nil {
		t.Errorf("ListRegions returned error: %v", err)
	}
	testResponseData(t, result, want)
}
