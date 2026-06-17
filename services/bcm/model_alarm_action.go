package bcm

type AlarmAction struct {
	Name          *string   `json:"name,omitempty"`
	BcmType       *string   `json:"type,omitempty"`
	ExecutedTime  *int64    `json:"executedTime,omitempty"`
	Notifications []*string `json:"notifications,omitempty"`
	CallBacks     []*string `json:"callBacks,omitempty"`
	Members       []*string `json:"members,omitempty"`
}
