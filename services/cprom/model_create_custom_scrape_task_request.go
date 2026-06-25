package cprom

type CreateCustomScrapeTaskRequest struct {
	InstanceId *string `json:"-"`
	AgentId    *string `json:"-"`
	Config     *string `json:"config,omitempty"`
}
