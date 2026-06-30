package bci

type HTTPGetAction struct {
	Path        *string       `json:"path,omitempty"`
	Port        *int32        `json:"port,omitempty"`
	Scheme      *string       `json:"scheme,omitempty"`
	Host        *string       `json:"host,omitempty"`
	HttpHeaders []*HTTPHeader `json:"httpHeaders,omitempty"`
}
