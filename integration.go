package runscope

import (
	"fmt"
	"net/http"
)

type Integration struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}

func (client *Client) ListIntegrations(teamID string) (*[]Integration, *http.Response, error) {
	var integrations = []Integration{}
	path := fmt.Sprintf("teams/%s/integrations", teamID)
	resp, err := client.Get(path, &integrations)
	return &integrations, resp, err
}
