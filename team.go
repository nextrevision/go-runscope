package runscope

import (
	"fmt"
	"net/http"
)

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

// TeamIntegration is a duplicate of Integration, except
// Runscope uses "integration_type" here instead of "type"
type TeamIntegration struct {
	Description string `json:"description"`
	Type        string `json:"integration_type"`
	ID          string `json:"id"`
	UUID        string `json:"uuid"`
}

func (client *Client) ListPeople(teamID string) (*[]Person, *http.Response, error) {
	var people = []Person{}
	path := fmt.Sprintf("teams/%s/people", teamID)
	resp, err := client.Get(path, &people)
	return &people, resp, err
}
