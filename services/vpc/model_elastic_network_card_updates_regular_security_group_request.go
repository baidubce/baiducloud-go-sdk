package vpc

type ElasticNetworkCardUpdatesRegularSecurityGroupRequest struct {
	EniId            *string   `json:"-"`
	ClientToken      *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
