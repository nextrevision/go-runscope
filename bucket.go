package runscope

import "net/http"

// Bucket struct
type Bucket struct {
	Name      string `json:"name"`
	AuthToken string `json:"auth_token"`
	Default   bool   `json:"default"`
	Key       string `json:"key"`
	Team      Team   `json:"team"`
	VerifySSL bool   `json:"verify_ssl"`
}

// NewBucketRequest struct
type NewBucketRequest struct {
	Name     string `json:"name"`
	TeamUUID string `json:"team_uuid"`
}

// ListBuckets returns a listing of all buckets
func (client *Client) ListBuckets() ([]Bucket, *http.Response, error) {
	var buckets = []Bucket{}
	resp, _, err := client.Get("buckets", &buckets)
	if err != nil {
		println(err.Error())
	}
	return buckets, resp, err
}

// GetBucket returns a listing of all buckets
func (client *Client) GetBucket(key string) (*Bucket, *http.Response, error) {
	var bucket = Bucket{}
	resp, _, err := client.Get("buckets/"+key, &bucket)
	if err != nil {
		println(err.Error())
	}
	return &bucket, resp, err
}

// NewBucket func
func (client *Client) NewBucket(name string, team Team) (*Bucket, *http.Response, error) {
	var bucket = Bucket{}
	data := NewBucketRequest{
		Name:     name,
		TeamUUID: team.UUID,
	}
	resp, _, err := client.Post("buckets", &data, &bucket)
	if err != nil {
		println(err.Error())
	}
	return &bucket, resp, err
}

// DeleteBucket func
func (client *Client) DeleteBucket(key string) (*http.Response, error) {
	resp, err := client.Delete("buckets/" + key)
	if err != nil {
		println(err.Error())
	}
	return resp, err
}
