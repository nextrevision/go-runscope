package runscope

import (
	"encoding/json"
	"fmt"
)

type Bucket struct {
	Name      string `json:"name"`
	AuthToken string `json:"auth_token"`
	Default   bool   `json:"default"`
	Key       string `json:"key"`
	Team      Team   `json:"team"`
	VerifySSL bool   `json:"verify_ssl"`
}

type NewBucketRequest struct {
	Name     string `json:"name"`
	TeamUUID string `json:"team_uuid"`
}

func (client *Client) ListBuckets() (*[]Bucket, error) {
	var buckets = []Bucket{}

	content, err := client.Get("buckets")
	if err != nil {
		return &buckets, err
	}

	err = unmarshal(content, &buckets)
	return &buckets, err
}

func (client *Client) GetBucket(bucketKey string) (*Bucket, error) {
	var bucket = Bucket{}

	path := fmt.Sprintf("buckets/%s", bucketKey)
	content, err := client.Get(path)
	if err != nil {
		return &bucket, err
	}

	err = unmarshal(content, &bucket)
	return &bucket, err
}

func (client *Client) NewBucket(newBucketRequest *NewBucketRequest) (*Bucket, error) {
	var bucket = Bucket{}

	data, err := json.Marshal(newBucketRequest)
	if err != nil {
		return &bucket, err
	}

	content, err := client.Post("buckets", data)
	if err != nil {
		return &bucket, err
	}

	err = unmarshal(content, &bucket)
	return &bucket, err
}

func (client *Client) DeleteBucket(bucketKey string) error {
	path := fmt.Sprintf("buckets/%s", bucketKey)
	return client.Delete(path)
}
