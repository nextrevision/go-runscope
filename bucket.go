package runscope

import (
	"fmt"
	"net/http"
)

type Bucket struct {
	Name      string `json:"name"`
	AuthToken string `json:"auth_token"`
	Default   bool   `json:"default"`
	Key       string `json:"key"`
	Team      Team   `json:"team"`
	VerifySSL bool   `json:"verify_ssl"`
}

type newBucketRequest struct {
	Name     string `json:"name"`
	TeamUUID string `json:"team_uuid"`
}

func (client *Client) ListBuckets() (*[]Bucket, *http.Response, error) {
	var buckets = []Bucket{}
	path := "buckets"
	resp, err := client.Get(path, &buckets)
	return &buckets, resp, err
}

func (client *Client) GetBucket(bucketKey string) (*Bucket, *http.Response, error) {
	var bucket = Bucket{}
	path := fmt.Sprintf("buckets/%s", bucketKey)
	resp, err := client.Get(path, &bucket)
	return &bucket, resp, err
}

func (client *Client) NewBucket(name string, teamUUID string) (*Bucket, *http.Response, error) {
	var newBucket = Bucket{}
	req := newBucketRequest{
		Name:     name,
		TeamUUID: teamUUID,
	}
	path := "buckets"
	resp, err := client.Post(path, &req, &newBucket)
	return &newBucket, resp, err
}

func (client *Client) DeleteBucket(bucketKey string) (*http.Response, error) {
	path := fmt.Sprintf("buckets/%s", bucketKey)
	resp, err := client.Delete(path)
	return resp, err
}
