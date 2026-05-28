package blb

type DescribeBlbTcpListenerRequest struct {
	BlbId        *string `json:"-"`
	ListenerPort *int32  `json:"-"`
	Marker       *string `json:"-"`
	MaxKeys      *int32  `json:"-"`
}
