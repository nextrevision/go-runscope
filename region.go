package runscope

type Regions struct {
	Regions []Region `json:"regions"`
}

type Region struct {
	RegionCode      string `json:"region_code"`
	Location        string `json:"location"`
	ServiceProvider string `json:"service_provider"`
	Hostname        string `json:"hostname"`
}

func (client *Client) ListRegions() (*Regions, error) {
	var regions = Regions{}

	content, err := client.Get("regions")
	if err != nil {
		return &regions, err
	}

	err = unmarshal(content, &regions)
	return &regions, err
}
