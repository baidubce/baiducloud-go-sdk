package eip

type TbspIpCleanModel struct {
	Ip               *string `json:"ip,omitempty"`
	EipName          *string `json:"eipName,omitempty"`
	EipId            *string `json:"eipId,omitempty"`
	ThresholdType    *string `json:"thresholdType,omitempty"`
	IpCleanMbps      *int32  `json:"ipCleanMbps,omitempty"`
	IpCleanPps       *int32  `json:"ipCleanPps,omitempty"`
	ProtectLevel     *string `json:"protectLevel,omitempty"`
	TurnOffBeginTime *string `json:"turnOffBeginTime,omitempty"`
	TurnOffEndTime   *string `json:"turnOffEndTime,omitempty"`
}
