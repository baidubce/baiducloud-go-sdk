package vpc

type ElasticNetworkCardUninstallationCloudProductInstanceRequest struct {
	EniId       *string `json:"-"`
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
}
