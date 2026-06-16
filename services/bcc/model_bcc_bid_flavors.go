package bcc

type BccBidFlavors struct {
	SpecId             *string `json:"specId,omitempty"`
	CpuCount           *int32  `json:"cpuCount,omitempty"`
	MemoryCapacityInGB *int32  `json:"memoryCapacityInGB,omitempty"`
	ProductType        *string `json:"productType,omitempty"`
	Spec               *string `json:"spec,omitempty"`
}
