package runscope

import "net/http"

type Account struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	ID        string  `json:"id"`
	UUID      string  `json:"uuid"`
	CreatedAt float64 `json:"created_at"`
	Teams     []Team  `json:"teams"`
}

func (client *Client) GetAccount() (*Account, *http.Response, error) {
	var account = Account{}
	path := "account"
	resp, err := client.Get(path, &account)
	return &account, resp, err
}
