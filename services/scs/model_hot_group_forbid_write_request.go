package scs

type HotGroupForbidWriteRequest struct {
	GroupId         *string `json:"-"`
	ForbidWriteFlag *bool   `json:"forbidWriteFlag,omitempty"`
}
