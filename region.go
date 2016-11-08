package runscope

// Regions represents multiple Runscope regions
type Regions struct {
	Regions []Region `json:"regions"`
}

// Region represents a Runscope region
type Region struct {
	RegionCode      string `json:"region_code"`
	Location        string `json:"location"`
	ServiceProvider string `json:"service_provider"`
	Hostname        string `json:"hostname"`
}

// ListRegions returns all regions known by Runscope
func (client *Client) ListRegions() (Regions, error) {
	var regions = Regions{}

	content, err := client.Get("regions")
	if err != nil {
		return regions, err
	}

	err = unmarshal(content, &regions)
	return regions, err
}
