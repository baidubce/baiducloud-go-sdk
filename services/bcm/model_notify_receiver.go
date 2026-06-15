package bcm

type NotifyReceiver struct {
	Id           *string   `json:"id,omitempty"`
	Name         *string   `json:"name,omitempty"`
	ReceiverType *string   `json:"receiverType,omitempty"`
	Channels     []*string `json:"channels,omitempty"`
}
