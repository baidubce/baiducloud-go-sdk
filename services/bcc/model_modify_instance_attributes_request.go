package bcc

type ModifyInstanceAttributesRequest struct {
	InstanceId       *string `json:"-"`
	Name             *string `json:"name,omitempty"`
	EnableJumboFrame *bool   `json:"enableJumboFrame,omitempty"`
	NetEthQueueCount *string `json:"netEthQueueCount,omitempty"`
}
