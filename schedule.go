package runscope

type Schedule struct {
	ID            string `json:"id"`
	Note          string `json:"note"`
	Interval      string `json:"interval"`
	EnvironmentID string `json:"environment_id"`
}
