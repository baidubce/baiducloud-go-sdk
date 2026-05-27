package vpc

type UpdateProbeRequest struct {
	ProbeId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DestIp      *string `json:"destIp,omitempty"`
	DestPort    *int32  `json:"destPort,omitempty"`
}
