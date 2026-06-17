package bcc

type OperationProgressSet struct {
	ResourceId      *string `json:"resourceId,omitempty"`
	OperationStatus *string `json:"operationStatus,omitempty"`
	Code            *string `json:"code,omitempty"`
	ErrorMessage    *string `json:"errorMessage,omitempty"`
}
