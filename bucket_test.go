package runscope

import (
	"net/http"
	"testing"
)

func TestListBuckets(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets"
	responseCode := http.StatusOK
	responseData := `
{
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
}`
	want := &[]Bucket{
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

	handleGet(t, path, responseCode, responseData)

	result, err := client.ListBuckets()
	if err != nil {
		t.Errorf("ListBuckets returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestGetBucket(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1"
	responseCode := http.StatusOK
	responseData := `
{
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
}`
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

	handleGet(t, path, responseCode, responseData)

	result, err := client.GetBucket("1")
	if err != nil {
		t.Errorf("GetBucket returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestNewBucket(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets"
	request := &NewBucketRequest{
		Name:     "Mobile Apps",
		TeamUUID: "7a7a0917-91d7-43ef-b8f4-fe31762167e0",
	}
	responseCode := http.StatusCreated
	responseData := `
{
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
}`
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

	handlePost(t, path, responseCode, responseData, new(NewBucketRequest), request)

	result, err := client.NewBucket(request)
	if err != nil {
		t.Errorf("NewBucket returned error: %v", err)
	}
	testResponseData(t, result, want)
}

func TestDeleteBucket(t *testing.T) {
	setup()
	defer teardown()

	path := "/buckets/1"
	responseCode := http.StatusNoContent

	handleDelete(t, path, responseCode)

	err := client.DeleteBucket("1")
	if err != nil {
		t.Errorf("DeleteBucket returned error: %v", err)
	}
}
