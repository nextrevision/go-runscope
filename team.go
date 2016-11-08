package runscope

import "fmt"

// Team represents a Runscope team
type Team struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	UUID string `json:"uuid"`
}

// Person can't be defined by your labels
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

// ListPeople returns a listing of people
func (client *Client) ListPeople(teamID string) ([]Person, error) {
	var people = []Person{}

	path := fmt.Sprintf("teams/%s/people", teamID)
	content, err := client.Get(path)
	if err != nil {
		return people, err
	}

	err = unmarshal(content, &people)
	return people, err
}
