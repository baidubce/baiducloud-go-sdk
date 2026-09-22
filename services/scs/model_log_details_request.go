package scs

type LogDetailsRequest struct {
	InstanceId   *string `json:"-"`
	LogId        *string `json:"-"`
	ValidSeconds *string `json:"-"`
}
