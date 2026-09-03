package image

type RetouchingResult struct {
	TaskId       *string `json:"task_id,omitempty"`
	Status       *string `json:"status,omitempty"`
	Dlink        *string `json:"dlink,omitempty"`
	CallbackData *string `json:"callback_data,omitempty"`
}
