package runscope

type Team struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	UUID string `json:"uuid"`
}

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
	UUID  string `json:"uuid"`
}

// TeamIntegration is a duplicate of Integration, except
// Runscope uses "integration_type" here instead of "type"
type TeamIntegration struct {
	Description string `json:"description"`
	Type        string `json:"integration_type"`
	ID          string `json:"id"`
	UUID        string `json:"uuid"`
}
