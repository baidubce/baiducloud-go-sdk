package bci

type Probe struct {
	InitialDelaySeconds           *int32           `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds                *int32           `json:"timeoutSeconds,omitempty"`
	PeriodSeconds                 *int32           `json:"periodSeconds,omitempty"`
	SuccessThreshold              *int32           `json:"successThreshold,omitempty"`
	FailureThreshold              *int32           `json:"failureThreshold,omitempty"`
	TerminationGracePeriodSeconds *int64           `json:"terminationGracePeriodSeconds,omitempty"`
	HttpGet                       *HTTPGetAction   `json:"httpGet,omitempty"`
	TcpSocket                     *TCPSocketAction `json:"tcpSocket,omitempty"`
	Exec                          *ExecAction      `json:"exec,omitempty"`
	Grpc                          *GRPCAction      `json:"grpc,omitempty"`
}
