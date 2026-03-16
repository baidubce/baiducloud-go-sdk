

package vpc

type CreateVpnRequest struct {
    ClientToken *string `json:"-"`
    VpcId *string `json:"vpcId,omitempty"`
    SubnetId *string `json:"subnetId,omitempty"`
    VpnName *string `json:"vpnName,omitempty"`
    Type *string `json:"type,omitempty"`
    Description *string `json:"description,omitempty"`
    Eip *string `json:"eip,omitempty"`
    Tags []*TagModel `json:"tags,omitempty"`
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`
    Billing *Billing `json:"billing,omitempty"`
    MaxConnection *int32 `json:"maxConnection,omitempty"`
    DeleteProtect *bool `json:"deleteProtect,omitempty"`
}