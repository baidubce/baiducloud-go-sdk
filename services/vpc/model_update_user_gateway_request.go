

package vpc

type UpdateUserGatewayRequest struct {
    CgwId *string `json:"-"`
    ClientToken *string `json:"-"`
    Name *string `json:"name,omitempty"`
    Description *string `json:"description,omitempty"`
}