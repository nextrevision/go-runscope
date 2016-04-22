package runscope

// Team struct
type Team struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// Person struct
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
	UUID  string `json:"uuid"`
}

// Integration Struct
type Integration struct {
	Description string `json:"description"`
	Type        string `json:"integration_type"`
	ID          string `json:"id"`
	UUID        string `json:"uuid"`
}
