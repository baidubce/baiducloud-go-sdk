package eip

type ResizeTbspRequest struct {
	Id          *string `json:"-"`
	ClientToken *string `json:"-"`
	IpCapacity  *int32  `json:"ipCapacity,omitempty"`
}
