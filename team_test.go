package runscope

import (
	"net/http"
	"testing"
)

func TestListPeople(t *testing.T) {
	setup()
	defer teardown()

	path := "/teams/1/people"
	responseCode := http.StatusOK
	responseData := `
{
  "data": [
    {
      "email": "grace@example.com",
      "name": "Grace Hopper",
      "uuid": "cf95026e-8951-4ae1-83a7-699243678490"
    }
  ],
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := []Person{
		Person{
			UUID:  "cf95026e-8951-4ae1-83a7-699243678490",
			Name:  "Grace Hopper",
			Email: "grace@example.com",
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, err := client.ListPeople("1")
	if err != nil {
		t.Errorf("ListPeople returned error: %v", err)
	}
	testResponseData(t, result, want)
}
