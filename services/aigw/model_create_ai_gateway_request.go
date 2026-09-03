package aigw

type CreateAIGatewayRequest struct {
	XRegion          *string        `json:"-"`
	Name             *string        `json:"name,omitempty"`
	VpcId            *string        `json:"vpcId,omitempty"`
	VpcCidr          *string        `json:"vpcCidr,omitempty"`
	SubnetId         *string        `json:"subnetId,omitempty"`
	GatewayType      *string        `json:"gatewayType,omitempty"`
	IsInternal       *string        `json:"isInternal,omitempty"`
	NetworkTypes     []*string      `json:"networkTypes,omitempty"`
	Replicas         *int32         `json:"replicas,omitempty"`
	InstallMode      *string        `json:"installMode,omitempty"`
	Description      *string        `json:"description,omitempty"`
	DeleteProtection *bool          `json:"deleteProtection,omitempty"`
	SrcProduct       *string        `json:"srcProduct,omitempty"`
	AccountId        *string        `json:"accountId,omitempty"`
	WorkspaceId      *string        `json:"workspaceId,omitempty"`
	WorkspaceName    *string        `json:"workspaceName,omitempty"`
	BlbId            *string        `json:"blbId,omitempty"`
	BlbIp            *string        `json:"blbIp,omitempty"`
	Clusters         []*ClusterInfo `json:"clusters,omitempty"`
	CpromInstanceId  *string        `json:"cpromInstanceId,omitempty"`
	CpromBearerToken *string        `json:"cpromBearerToken,omitempty"`
	BlsEnabled       *bool          `json:"blsEnabled,omitempty"`
	LogStoreName     *string        `json:"logStoreName,omitempty"`
	Version          *string        `json:"version,omitempty"`
	Tags             []*Tag         `json:"tags,omitempty"`
	ResourceGroupId  *string        `json:"resourceGroupId,omitempty"`
	AihcArgs         *AihcArgs      `json:"aihcArgs,omitempty"`
}
