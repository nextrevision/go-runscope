package runscope

import "net/http"

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
	resp, _, err := client.Get("buckets", &buckets)
	if err != nil {
		println(err.Error())
	}
	return &buckets, resp, err
}

func (client *Client) GetBucket(key string) (*Bucket, *http.Response, error) {
	var bucket = Bucket{}
	resp, _, err := client.Get("buckets/"+key, &bucket)
	if err != nil {
		println(err.Error())
	}
	return &bucket, resp, err
}

func (client *Client) NewBucket(name string, teamUUID string) (*Bucket, *http.Response, error) {
	var newBucket = Bucket{}
	req := newBucketRequest{
		Name:     name,
		TeamUUID: teamUUID,
	}
	resp, _, err := client.Post("buckets", &req, &newBucket)
	if err != nil {
		println(err.Error())
	}
	return &newBucket, resp, err
}

func (client *Client) DeleteBucket(key string) (*http.Response, error) {
	resp, err := client.Delete("buckets/" + key)
	if err != nil {
		println(err.Error())
	}
	return resp, err
}
