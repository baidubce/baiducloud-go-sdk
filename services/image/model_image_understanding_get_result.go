package image

type ImageUnderstandingGetResult struct {
	TaskId      *string `json:"task_id,omitempty"`
	RetCode     *int32  `json:"ret_code,omitempty"`
	RetMsg      *string `json:"ret_msg,omitempty"`
	Description *string `json:"description,omitempty"`
}
