package scs

type SyncGroupModifyBnsgroupRequest struct {
	GroupId  *string `json:"-"`
	BnsGroup *string `json:"bnsGroup,omitempty"`
}
