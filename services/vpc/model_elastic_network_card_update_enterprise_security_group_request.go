package vpc

type ElasticNetworkCardUpdateEnterpriseSecurityGroupRequest struct {
	EniId                      *string   `json:"-"`
	ClientToken                *string   `json:"-"`
	EnterpriseSecurityGroupIds []*string `json:"enterpriseSecurityGroupIds,omitempty"`
}
