package cprom

type DeleteCustomScrapeTaskRequest struct {
	ScrapeJobId *string `json:"-"`
	InstanceId  *string `json:"-"`
	AgentId     *string `json:"-"`
}
