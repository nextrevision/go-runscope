package runscope

// Account struct
type Account struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	UUID  string `json:"uuid"`
	Teams []Team `json:"teams"`
}
