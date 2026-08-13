package aigw

type TargetService struct {
	ServiceSource        *string `json:"serviceSource,omitempty"`
	ServiceName          *string `json:"serviceName,omitempty"`
	Namespace            *string `json:"namespace,omitempty"`
	ServicePort          *int32  `json:"servicePort,omitempty"`
	LoadBalanceAlgorithm *string `json:"loadBalanceAlgorithm,omitempty"`
	HashType             *string `json:"hashType,omitempty"`
	HashKey              *string `json:"hashKey,omitempty"`
	RequestRatio         *int32  `json:"requestRatio,omitempty"`
	WeightFactor         *int32  `json:"weightFactor,omitempty"`
	ModelName            *string `json:"modelName,omitempty"`
	ModelNameMode        *string `json:"modelNameMode,omitempty"`
	SpecifiedModelName   *string `json:"specifiedModelName,omitempty"`
}
