package bcc

type UnbindReservedInstanceFromTagsRequest struct {
	ReservedInstanceIds []*string   `json:"reservedInstanceIds,omitempty"`
	ChangeTags          []*TagModel `json:"changeTags,omitempty"`
}
