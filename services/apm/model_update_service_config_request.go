package apm

type UpdateServiceConfigRequest struct {
	ServiceNames           []*string               `json:"serviceNames,omitempty"`
	SampleConfig           *SampleConfig           `json:"sampleConfig,omitempty"`
	LoggingConfig          *LoggingConfig          `json:"loggingConfig,omitempty"`
	RequestConfig          *ServiceRequestConfig   `json:"requestConfig,omitempty"`
	TopoConfig             *ServiceTopoConfig      `json:"topoConfig,omitempty"`
	MllmResourceDumpConfig *MllmResourceDumpConfig `json:"mllmResourceDumpConfig,omitempty"`
}
