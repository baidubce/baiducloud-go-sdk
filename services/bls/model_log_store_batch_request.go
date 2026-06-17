package bls

type LogStoreBatchRequest struct {
	Project      *string `json:"project,omitempty"`
	LogStoreName *string `json:"logStoreName,omitempty"`
}
