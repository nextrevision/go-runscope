package runscope

import "fmt"

type Integration struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}

func (client *Client) ListIntegrations(teamID string) (*[]Integration, error) {
	var integrations = []Integration{}

	path := fmt.Sprintf("teams/%s/integrations", teamID)
	content, err := client.Get(path)
	if err != nil {
		return &integrations, err
	}

	err = unmarshal(content, &integrations)
	return &integrations, err
}
