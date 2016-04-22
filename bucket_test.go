package runscope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListBuckets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
          "data": [
              {
                  "auth_token": null,
                  "default": false,
                  "key": "z20co0kgljjk",
                  "name": "Lucky Notebook",
                  "team": {
                      "name": "Personal Team",
                      "uuid": "7a7a0917-91d7-43ef-b8f4-fe31762167e0"
                  },
                  "verify_ssl": true
              },
              {
                  "auth_token": null,
                  "default": false,
                  "key": "ov2f2tqifoov",
                  "auth_token": "7n7n0917-91q7-43rs-o8s4-sr31762167r0",
                  "name": "Mobile Apps",
                  "team": {
                      "name": "Mobile Team",
                      "uuid": "7a7a0917-91d7-43ef-b8f4-fe31762167e0"
                  },
                  "verify_ssl": true
              }
          ],
          "meta": {
              "status": "success"
          }
      }`)
	})

	buckets, _, err := client.ListBuckets()
	if err != nil {
		t.Errorf("ListBuckets returned error: %v", err)
	}

	want := []Bucket{
		Bucket{
			Name:      "Lucky Notebook",
			Default:   false,
			Key:       "z20co0kgljjk",
			AuthToken: "",
			VerifySSL: true,
			Team: Team{
				Name: "Personal Team",
				UUID: "7a7a0917-91d7-43ef-b8f4-fe31762167e0",
			},
		},
		Bucket{
			Name:      "Mobile Apps",
			Default:   false,
			Key:       "ov2f2tqifoov",
			VerifySSL: true,
			AuthToken: "7n7n0917-91q7-43rs-o8s4-sr31762167r0",
			Team: Team{
				Name: "Mobile Team",
				UUID: "7a7a0917-91d7-43ef-b8f4-fe31762167e0",
			},
		},
	}
	if !reflect.DeepEqual(buckets, want) {
		t.Errorf("ListBuckets returned %+v, want %+v", buckets, want)
	}
}

func TestGetBucket(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
         "data": {
            "auth_token": null,
            "default": false,
            "key": "ov2f2tq1floq",
            "name": "Mobile Apps",
            "team": {
              "name": "Mobile Team",
              "uuid": "7a7a0917-91d7-43ef-b8f4-fe31762167e0"
            },
            "verify_ssl": true
        },
        "meta": {
          "status": "success"
        }
      }`)
	})

	bucket, _, err := client.GetBucket("abcde12345")
	if err != nil {
		t.Errorf("GetBucket returned error: %v", err)
	}

	want := &Bucket{
		Name:      "Mobile Apps",
		Default:   false,
		Key:       "ov2f2tq1floq",
		VerifySSL: true,
		AuthToken: "",
		Team: Team{
			Name: "Mobile Team",
			UUID: "7a7a0917-91d7-43ef-b8f4-fe31762167e0",
		},
	}
	if !reflect.DeepEqual(bucket, want) {
		t.Errorf("GetBucket returned %+v, want %+v", bucket, want)
	}
}

func TestNewBucket(t *testing.T) {
	setup()
	defer teardown()

	req := &newBucketRequest{
		Name:     "Mobile Apps",
		TeamUUID: "7a7a0917-91d7-43ef-b8f4-fe31762167e0",
	}

	mux.HandleFunc("/buckets", func(w http.ResponseWriter, r *http.Request) {
		v := new(newBucketRequest)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, req) {
			t.Errorf("Request body = %+v, want %+v", v, req)
		}
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,
			`{
         "data": {
            "auth_token": null,
            "default": false,
            "key": "ov2f2tq1floq",
            "name": "Mobile Apps",
            "team": {
              "name": "Mobile Team",
              "uuid": "7a7a0917-91d7-43ef-b8f4-fe31762167e0"
            },
            "verify_ssl": true
        },
        "meta": {
          "status": "success"
        }
      }`)
	})

	bucket, resp, err := client.NewBucket("Mobile Apps", "7a7a0917-91d7-43ef-b8f4-fe31762167e0")
	if err != nil {
		t.Errorf("NewBucket returned error: %v", err)
	}

	want := &Bucket{
		Name:      "Mobile Apps",
		Default:   false,
		Key:       "ov2f2tq1floq",
		VerifySSL: true,
		AuthToken: "",
		Team: Team{
			Name: "Mobile Team",
			UUID: "7a7a0917-91d7-43ef-b8f4-fe31762167e0",
		},
	}
	if resp.StatusCode != 201 {
		t.Errorf("NewBucket did not return 201: %v", resp)
	}
	if !reflect.DeepEqual(bucket, want) {
		t.Errorf("NewBucket returned %+v, want %+v", bucket, want)
	}
}

func TestDeleteBucket(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/buckets/abcde12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.DeleteBucket("abcde12345")
	if resp.StatusCode != 204 {
		t.Errorf("DeleteBucket did not return 204: %v", resp)
	}
	if err != nil {
		t.Errorf("DeleteBucket returned error: %v", err)
	}
}
