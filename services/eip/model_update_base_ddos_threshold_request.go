package eip

type UpdateBaseDdosThresholdRequest struct {
	Ip            *string `json:"-"`
	ClientToken   *string `json:"-"`
	ThresholdType *string `json:"thresholdType,omitempty"`
	IpCleanMbps   *int64  `json:"ipCleanMbps,omitempty"`
	IpCleanPps    *int64  `json:"ipCleanPps,omitempty"`
}
