package vdb

type TargetPackage struct {
	TargetMilvusKernelVersion *string `json:"targetMilvusKernelVersion,omitempty"`
	TargetPackageVersion      *string `json:"targetPackageVersion,omitempty"`
	UpgradeNote               *string `json:"upgradeNote,omitempty"`
}
