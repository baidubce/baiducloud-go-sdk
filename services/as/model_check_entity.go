package as

type CheckEntity struct {
	Label  *string      `json:"label,omitempty"`
	Status *interface{} `json:"status,omitempty"`
	Result *CheckDetail `json:"result,omitempty"`
}
