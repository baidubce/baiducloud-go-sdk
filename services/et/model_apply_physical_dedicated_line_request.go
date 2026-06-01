package et

type ApplyPhysicalDedicatedLineRequest struct {
	ClientToken *string      `json:"-"`
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Isp         *string      `json:"isp,omitempty"`
	IntfType    *string      `json:"intfType,omitempty"`
	ApType      *string      `json:"apType,omitempty"`
	ApAddr      *string      `json:"apAddr,omitempty"`
	LinkDelay   *int32       `json:"linkDelay,omitempty"`
	UserName    *string      `json:"userName,omitempty"`
	UserPhone   *string      `json:"userPhone,omitempty"`
	UserEmail   *string      `json:"userEmail,omitempty"`
	UserIdc     *string      `json:"userIdc,omitempty"`
	Billing     *Billing     `json:"billing,omitempty"`
	AutoRenew   *Reservation `json:"autoRenew,omitempty"`
	Tags        []*TagModel  `json:"tags,omitempty"`
}
