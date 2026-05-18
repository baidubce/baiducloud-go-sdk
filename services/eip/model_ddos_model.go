package eip

type DdosModel struct {
	Ip               *string `json:"ip,omitempty"`
	Status           *string `json:"status,omitempty"`
	BindInstanceType *string `json:"bindInstanceType,omitempty"`
	BindInstanceId   *string `json:"bindInstanceId,omitempty"`
	IpCleanMbps      *int64  `json:"ipCleanMbps,omitempty"`
	IpCleanPps       *int64  `json:"ipCleanPps,omitempty"`
	ThresholdType    *string `json:"thresholdType,omitempty"`
	MaximumThreshold *int64  `json:"maximumThreshold,omitempty"`
}
