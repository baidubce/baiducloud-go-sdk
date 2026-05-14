package snic

type UpdateSnicSgRequest struct {
	EndpointId       *string   `json:"-"`
	ClientToken      *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
