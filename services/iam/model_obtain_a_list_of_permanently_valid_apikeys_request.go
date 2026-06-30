package iam

type ObtainAListOfPermanentlyValidApikeysRequest struct {
	UserId   *string   `json:"userId,omitempty"`
	Service  []*string `json:"service,omitempty"`
	PageNo   *int32    `json:"pageNo,omitempty"`
	PageSize *int32    `json:"pageSize,omitempty"`
}
