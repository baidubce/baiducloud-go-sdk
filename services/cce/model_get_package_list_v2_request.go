package cce

type GetPackageListV2Request struct {
	Type            *string   `json:"-"`
	MachineSpecList []*string `json:"machineSpecList,omitempty"`
}
