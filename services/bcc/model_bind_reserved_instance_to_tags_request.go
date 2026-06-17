package bcc

type BindReservedInstanceToTagsRequest struct {
	ReservedInstanceIds []*string   `json:"reservedInstanceIds,omitempty"`
	ChangeTags          []*TagModel `json:"changeTags,omitempty"`
}
