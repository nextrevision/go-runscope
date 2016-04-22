package runscope

type Team struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
	UUID  string `json:"uuid"`
}

type Integration struct {
	Description string `json:"description"`
	Type        string `json:"integration_type"`
	ID          string `json:"id"`
	UUID        string `json:"uuid"`
}
