package runscope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListSchedules(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345/tests/12345/schedules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
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
			}`)
	})

	schedules, _, err := client.ListSchedules("abcde12345", "12345")
	if err != nil {
		t.Errorf("ListSchedules returned error: %v", err)
	}

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
	if !reflect.DeepEqual(schedules, want) {
		t.Errorf("ListSchedules returned %+v, want %+v", schedules, want)
	}
}

func TestGetSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345/tests/12345/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
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
      }`)
	})

	schedule, _, err := client.GetSchedule("abcde12345", "12345", "1")
	if err != nil {
		t.Errorf("GetSchedule returned error: %v", err)
	}

	want := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		Note:          "Staging Environment",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "1h",
	}
	if !reflect.DeepEqual(schedule, want) {
		t.Errorf("GetSchedule returned %+v, want %+v", schedule, want)
	}
}

func TestNewSchedule(t *testing.T) {
	setup()
	defer teardown()

	req := &Schedule{
		Interval:      "6h",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
	}

	mux.HandleFunc("/buckets/abcde12345/tests/12345/schedules", func(w http.ResponseWriter, r *http.Request) {
		v := new(Schedule)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,
			`{
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
      }`)
	})

	schedule, resp, err := client.NewSchedule("abcde12345", "12345", req)
	if err != nil {
		t.Errorf("NewSchedule returned error: %v", err)
	}

	want := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "6h",
	}
	if resp.StatusCode != 201 {
		t.Errorf("NewSchedule did not return 201: %v", resp)
	}
	if !reflect.DeepEqual(schedule, want) {
		t.Errorf("NewSchedule returned %+v, want %+v", schedule, want)
	}
}

func TestUpdateSchedule(t *testing.T) {
	setup()
	defer teardown()

	req := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "12h",
	}

	mux.HandleFunc("/buckets/abcde12345/tests/12345/schedules/084e6df7-9165-46d2-9e1c-b87ccfc53d18", func(w http.ResponseWriter, r *http.Request) {
		v := new(Schedule)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
         "data": {
					  "environment_id": "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
						"interval": "12h",
						"note": "",
						"id": "084e6df7-9165-46d2-9e1c-b87ccfc53d18"
        },
        "meta": {
          "status": "success"
        }
			}`)
	})

	schedule, resp, err := client.UpdateSchedule("abcde12345", "12345", req)
	if err != nil {
		t.Errorf("UpdateSchedule returned error: %v", err)
	}

	want := &Schedule{
		ID:            "084e6df7-9165-46d2-9e1c-b87ccfc53d18",
		EnvironmentID: "1eeb3695-5d0f-467c-9d51-8b773dce29ba",
		Interval:      "12h",
	}
	if resp.StatusCode != 200 {
		t.Errorf("UpdateSchedule did not return 200: %v", resp)
	}
	if !reflect.DeepEqual(schedule, want) {
		t.Errorf("UpdateSchedule returned %+v, want %+v", schedule, want)
	}
}

func TestDeleteSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345/tests/12345/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.DeleteSchedule("abcde12345", "12345", "1")
	if resp.StatusCode != 204 {
		t.Errorf("DeleteSchedule did not return 204: %v", resp)
	}
	if err != nil {
		t.Errorf("DeleteSchedule returned error: %v", err)
	}
}
