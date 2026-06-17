package bls

type UpdateFastQueryRequest struct {
	Name          *string `json:"-"`
	FastQueryName *string `json:"fastQueryName,omitempty"`
	Query         *string `json:"query,omitempty"`
	Description   *string `json:"description,omitempty"`
	Project       *string `json:"project,omitempty"`
	LogStoreName  *string `json:"logStoreName,omitempty"`
	LogStreamName *string `json:"logStreamName,omitempty"`
}
