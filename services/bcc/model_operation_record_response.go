package bcc

type OperationRecordResponse struct {
	Name        *string `json:"name,omitempty"`
	Operator    *string `json:"operator,omitempty"`
	OperateTime *string `json:"operateTime,omitempty"`
}
