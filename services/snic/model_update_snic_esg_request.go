package snic

type UpdateSnicEsgRequest struct {
	EndpointId                 *string   `json:"-"`
	ClientToken                *string   `json:"-"`
	EnterpriseSecurityGroupIds []*string `json:"enterpriseSecurityGroupIds,omitempty"`
}
