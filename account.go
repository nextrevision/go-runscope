package runscope

type Account struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	ID        string  `json:"id"`
	UUID      string  `json:"uuid"`
	CreatedAt float64 `json:"created_at"`
	Teams     []Team  `json:"teams"`
}

func (client *Client) GetAccount() (*Account, error) {
	var account = Account{}

	content, err := client.Get("account")
	if err != nil {
		return &account, err
	}

	err = unmarshal(content, &account)
	return &account, err
}
