package runscope

import (
	"net/http"
	"testing"
)

func TestListSchedules(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/schedules"
	responseCode := http.StatusOK
	responseData := `
{
  "data": [
    {
      "environment_id": "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
      "interval": "1h",
      "note": "Staging Environment",
      "id": "084e6df7-9165-46d2-9e1c-b87ccfc53d18"
    },
    {
      "environment_id": "5e70db57-7485-4ca9-bb4d-482416993ddd",
      "interval": "5m",
      "note": "Production Monitoring",
      "id": "c60e5a78-0dbd-493a-9e99-9a8282935d0c"
    }
  ],
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := &[]Schedule{
		Schedule{
			ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
			EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
			Interval:      "1h",
			Note:          "Staging Environment",
		},
		Schedule{
			ID:            "c60e5a78-0dbd-493a-9e99-9a8282935d0c",
			EnvironmentID: "5e70db57-7485-4ca9-bb4d-482416993ddd",
			Interval:      "5m",
			Note:          "Production Monitoring",
		},
	}

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.ListSchedules("1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("ListSchedules returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetSchedule(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/schedules/1"
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "environment_id": "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
    "interval": "1h",
    "note": "Staging Environment",
    "id": "084e6df7-9165-46d2-9e1c-b87ccfc53d18"
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		Note:          "Staging Environment",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "1h",
	}

	handleGet(t, path, responseCode, responseData)

	result, resp, err := client.GetSchedule("1", "1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("GetSchedule returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestNewSchedule(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/schedules"
	request := &Schedule{
		Interval:      "6h",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
	}
	responseCode := http.StatusCreated
	responseData := `
{
  "data": {
    "environment_id": "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
    "interval": "6h",
    "note": "",
    "id": "084e6df7-9165-46d2-9e1c-b87ccfc53d18"
  },
  "error": null,
  "meta": {
    "status": "success"
  }
}`
	want := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "6h",
	}

	handlePost(t, path, responseCode, responseData, new(Schedule), request)

	result, resp, err := client.NewSchedule("1", "1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("NewSchedule returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestUpdateSchedule(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/schedules/1"
	request := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "12h",
	}
	responseCode := http.StatusOK
	responseData := `
{
  "data": {
    "environment_id": "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
    "interval": "12h",
    "note": "",
    "id": "084e6df7-9165-46d2-9e1c-b87ccfc53d18"
  },
  "meta": {
    "status": "success"
  }
}`
	want := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "12h",
	}

	handlePut(t, path, responseCode, responseData, new(Schedule), request)

	result, resp, err := client.UpdateSchedule("1", "1", "1", request)
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("UpdateSchedule returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestDeleteSchedule(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1/tests/1/schedules/1"
	responseCode := http.StatusNoContent

	handleDelete(t, path, responseCode)

	resp, err := client.DeleteSchedule("1", "1", "1")
	testStatusCode(t, resp, responseCode)
	if err != nil {
		t.Errorf("DeleteSchedule returned error: %v", err)
	}
}
