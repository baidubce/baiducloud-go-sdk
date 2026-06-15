package bcc

type InstanceBatchResizeBySpecRequest struct {
	Spec             *string   `json:"spec,omitempty"`
	InstanceIdList   []*string `json:"instanceIdList,omitempty"`
	EnableJumboFrame *bool     `json:"enableJumboFrame,omitempty"`
	SubnetId         *string   `json:"subnetId,omitempty"`
	LogicalZone      *string   `json:"logicalZone,omitempty"`
}
