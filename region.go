package runscope

import "net/http"

type Regions struct {
	Regions []Region `json:"regions"`
}

type Region struct {
	RegionCode      string `json:"region_code"`
	Location        string `json:"location"`
	ServiceProvider string `json:"service_provider"`
	Hostname        string `json:"hostname"`
}

func (client *Client) ListRegions() (*Regions, *http.Response, error) {
	var regions = Regions{}
	path := "regions"
	resp, err := client.Get(path, &regions)
	return &regions, resp, err
}
