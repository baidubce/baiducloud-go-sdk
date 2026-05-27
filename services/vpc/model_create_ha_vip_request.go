package vpc

type CreateHaVipRequest struct {
	ClientToken      *string `json:"-"`
	Name             *string `json:"name,omitempty"`
	SubnetId         *string `json:"subnetId,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	Description      *string `json:"description,omitempty"`
}
