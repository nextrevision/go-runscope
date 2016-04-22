package runscope

type Environment struct {
	Name                string            `json:"name"`
	ID                  string            `json:"id"`
	ParentEnvironmentID string            `json:"parent_environment_id"`
	PreserveCookies     bool              `json:"preserve_cookies"`
	Regions             []string          `json:"regions"`
	RemoteAgents        []RemoteAgent     `json:"remote_agents"`
	Script              string            `json:"script"`
	TestID              string            `json:"test_id"`
	VerifySSL           bool              `json:"verify_ssl"`
	Webhooks            []string          `json:"webhooks"`
	Emails              Email             `json:"emails"`
	InitialVariables    map[string]string `json:"initial_variables"`
	Integrations        []Integration     `json:"integrations"`
}

type Email struct {
	NotifyAll       bool     `json:"notify_all"`
	NotifyOn        string   `json:"notify_on"`
	NotifyThreshold int      `json:"notify_threshold"`
	Recipients      []Person `json:"recipients"`
}

type RemoteAgent struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
