package bcc

type BccFlavor struct {
	CpuCount            *int32    `json:"cpuCount,omitempty"`
	MemoryCapacityInGB  *int32    `json:"memoryCapacityInGB,omitempty"`
	EphemeralDiskInGb   *int32    `json:"ephemeralDiskInGb,omitempty"`
	EphemeralDiskCount  *int32    `json:"ephemeralDiskCount,omitempty"`
	EphemeralDiskType   *string   `json:"ephemeralDiskType,omitempty"`
	GpuCardType         *string   `json:"gpuCardType,omitempty"`
	GpuCardCount        *int32    `json:"gpuCardCount,omitempty"`
	FpgaCardType        *string   `json:"fpgaCardType,omitempty"`
	FpgaCardCount       *int32    `json:"fpgaCardCount,omitempty"`
	ProductType         *string   `json:"productType,omitempty"`
	Spec                *string   `json:"spec,omitempty"`
	SpecId              *string   `json:"specId,omitempty"`
	EnableJumboFrame    *bool     `json:"enableJumboFrame,omitempty"`
	CpuModel            *string   `json:"cpuModel,omitempty"`
	CpuGHz              *string   `json:"cpuGHz,omitempty"`
	NetworkBandwidth    *string   `json:"networkBandwidth,omitempty"`
	NetworkPackage      *string   `json:"networkPackage,omitempty"`
	NetEthQueueCount    *string   `json:"netEthQueueCount,omitempty"`
	NetEthMaxQueueCount *string   `json:"netEthMaxQueueCount,omitempty"`
	EniQuota            *int32    `json:"eniQuota,omitempty"`
	EriQuota            *int32    `json:"eriQuota,omitempty"`
	RdmaType            *string   `json:"rdmaType,omitempty"`
	RdmaNetCardCount    *int32    `json:"rdmaNetCardCount,omitempty"`
	RdmaNetBandwidth    *int32    `json:"rdmaNetBandwidth,omitempty"`
	SystemDiskType      []*string `json:"systemDiskType,omitempty"`
	DataDiskType        []*string `json:"dataDiskType,omitempty"`
	NicIpv4Quota        *int32    `json:"nicIpv4Quota,omitempty"`
	NicIpv6Quota        *int32    `json:"nicIpv6Quota,omitempty"`
	VolumeCount         *int32    `json:"volumeCount,omitempty"`
}
