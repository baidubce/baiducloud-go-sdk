package ax

type ListSandboxesV2Request struct {
	Limit     *int32  `json:"-"`
	NextToken *string `json:"-"`
	Metadata  *string `json:"-"`
	State     *string `json:"-"`
}
