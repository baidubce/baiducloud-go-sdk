package cloudassistant

type ActionResult struct {
	ActionId   *string `json:"actionId,omitempty"`
	ActionName *string `json:"actionName,omitempty"`
	RunId      *string `json:"runId,omitempty"`
}
