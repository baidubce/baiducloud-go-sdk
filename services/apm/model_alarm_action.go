package apm

type AlarmAction struct {
	NotifyId     *string        `json:"notifyId,omitempty"`
	Alias        *string        `json:"alias,omitempty"`
	CallBacks    []*string      `json:"callBacks,omitempty"`
	DisableTimes []*DisableTime `json:"disableTimes,omitempty"`
	Members      []*string      `json:"members,omitempty"`
	Types        []*string      `json:"types,omitempty"`
}
