package aihc

type DescribeJobWebterminalRequest struct {
	ResourcePoolId         *string `json:"-"`
	QueueID                *string `json:"-"`
	JobId                  *string `json:"jobId,omitempty"`
	PodName                *string `json:"podName,omitempty"`
	HandshakeTimeoutSecond *string `json:"handshakeTimeoutSecond,omitempty"`
	PingTimeoutSecond      *string `json:"pingTimeoutSecond,omitempty"`
}
