package eip

type ModifyTbspIpCleanThresholdRequest struct {
	Id            *string `json:"-"`
	ClientToken   *string `json:"-"`
	Ip            *string `json:"ip,omitempty"`
	ThresholdType *string `json:"thresholdType,omitempty"`
	CleanMbps     *int32  `json:"cleanMbps,omitempty"`
	CleanPps      *int32  `json:"cleanPps,omitempty"`
}
